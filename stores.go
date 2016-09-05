package betwixt

type Store interface {
	Init()
	Close()

	GetClient(string) RegisteredClient
	GetClients() map[string]RegisteredClient
	PutClient(id string, c RegisteredClient)
	DeleteClient(id string)
	UpdateTS(id string)
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		connectedClients: make(map[string]RegisteredClient),
	}
}

type InMemoryStore struct {
	connectedClients map[string]RegisteredClient
}

func (db *InMemoryStore) Init() {

}

func (db *InMemoryStore) Close() {

}

func (db *InMemoryStore) GetClient(name string) RegisteredClient {
	return db.connectedClients[name]
}

func (db *InMemoryStore) GetClients() map[string]RegisteredClient {
	return db.connectedClients
}

func (db *InMemoryStore) PutClient(name string, c RegisteredClient) {
	db.connectedClients[name] = c
}

func (db *InMemoryStore) DeleteClient(name string) {
	for k, v := range db.connectedClients {
		if v.GetId() == name {

			delete(db.connectedClients, k)
			return
		}
	}
}

func (db *InMemoryStore) UpdateTS(name string) {
	for k, v := range db.connectedClients {
		if v.GetId() == name {
			v.Update()
			db.connectedClients[k] = v
		}
	}
}
