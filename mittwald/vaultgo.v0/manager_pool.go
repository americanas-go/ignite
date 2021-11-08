package vault

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/log"
	"github.com/hashicorp/vault/api"
	vault "github.com/mittwald/vaultgo"
)

type ManagerPool struct {
	managers []Manager
	client   *vault.Client
}

func NewManagerPool(client *vault.Client, managers ...Manager) *ManagerPool {
	return &ManagerPool{managers: managers, client: client}
}

func ManageAll(ctx context.Context, managers ...Manager) {
	client, err := NewClient(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	mp := NewManagerPool(client, managers...)
	mp.ManageAll(ctx)
}

func (m *ManagerPool) ManageAll(ctx context.Context) {

	for _, manager := range m.managers {
		mr := manager
		go func() {
			err := m.Configure(ctx, mr)
			if err != nil {
				log.Errorf("error on start vault manager. %s", err.Error())
			}
		}()
	}

}

func (m *ManagerPool) Configure(ctx context.Context, manager Manager) error {

	var response api.Secret

	err := m.client.Read([]string{manager.Options().SecretPath}, &response, &vault.RequestOptions{
		Parameters:  nil,
		SkipRenewal: false,
	})
	if err != nil {
		return err
	}

	log.Debugf("lease_id: %s", response.LeaseID)
	log.Debugf("data: %v", response.Data)
	log.Debugf("lease_duration: %vs", response.LeaseDuration)

	data := response.Data
	dataConv := make(map[string]interface{})

	options := manager.Options()

	for source, dst := range options.Keys {
		if dt, ok := data[source]; ok {
			dataConv[dst] = dt
		} else {
			log.Warnf("the key %s not found in vault data", source)
		}
	}

	if err := manager.Configure(ctx, dataConv); err != nil {
		return err
	}

	if manager.Options().Watcher.Enabled {
		go func() {
			err := m.watch(ctx, manager, response)
			if err != nil {
				log.Errorf("error on start vault watcher. %s", err.Error())
			}
		}()
	}

	return nil
}

func (m *ManagerPool) watch(ctx context.Context, manager Manager, response api.Secret) error {

	secretesTokenWatcher := api.LifetimeWatcherInput{
		Secret:    &response,
		Increment: manager.Options().Watcher.Increment,
	}
	watcher, err := m.client.NewLifetimeWatcher(&secretesTokenWatcher)
	if err != nil {
		return errors.Internalf("error on start watcher. %s", err.Error())
	}
	go watcher.Start()

	for {
		select {
		case rawData := <-watcher.RenewCh():
			log.Debugf("received renewal at: %+v", rawData.RenewedAt)
			log.Debugf("received renewal Secret: %+v", rawData.Secret)
		case er := <-watcher.DoneCh():
			if err != nil {
				log.Errorf("Got watcher error: %s", er.Error())
			}
			watcher.Stop()
			if err := manager.Close(ctx); err != nil {
				log.Errorf("Got manager error: %s", er.Error())
			}
			go func() {
				err := m.Configure(ctx, manager)
				if err != nil {
					log.Errorf("error on start vault manager. %s", err.Error())
				}
			}()
			return nil
		}
	}

}
