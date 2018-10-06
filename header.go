package goptrail

import "net"

type Client interface {
	ListUsers() ([]User, error)
	InviteUser(User) error
	UpdateUser(User) error
	DeleteUser(User) error

	ListLogDestinations() ([]LogDestination, error)
	RegisterSystem(s InputSystem) (OutputSystem, error)
	GetSystem(id string) (*OutputSystem, error)
	ListSystems() ([]OutputSystem, error)
	UpdateSystem(s InputSystem) error
	UnregisterSystem(id string) error
	AddSystemToGroup(sID, gID string) error
	RemoveSystemFromGroup(sID, gID string) error

	CreateGroup(g Group) (Group, error)
	GetGroup(id string) (Group, error)
	ListGroups() ([]Group, error)
	UpdateGroup(g Group) error
	DeleteGroup(id string) error

	CreateSearch(s Search) (Search, error)
	GetSearch(id string) (Search, error)
	ListSearch() ([]Search, error)
	UpdateSearch(s Search) error
	DeleteSearch(id string) error
}

type Search struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Query string `json:"query"`
	Group Group  `json:"group"`
}

type Group struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	SystemWildcard string         `json:"system_wildcard"`
	Systems        []OutputSystem `json:"systems"`
}

type SysLog struct {
	Hostname    string `json:"hostname"`
	Port        int    `json:"port"`
	Description string `json:"description"`
}

type LogDestination struct {
	ID     int    `json:"id"`
	Syslog SysLog `json:"syslog"`
}

type User struct {
	Email              string `json:"email"`
	ID                 int    `json:"id"`
	ReadOnly           int    `json:"read_only"`
	ManageMembers      int    `json:"manage_members"`
	ManageBilling      int    `json:"manage_billing"`
	PurgeLogs          int    `json:"purge_logs"`
	CanAccessAllGroups int    `json:"can_access_all_groups"`
	GroupIDs           []int  `json:"group_ids"`
}

type InputSystem struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	IpAddress       string `json:"ip_address"`
	Hostname        string `json:"hostname"`
	DestinationID   int    `json:"destination_id"`
	DestinationPort int    `json:"destination_port"`
	Description     string `json:"description"`
}

type OutputSystem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IpAddress   net.IP `json:"ip_address"`
	Hostname    string `json:"hostname"`
	LastEventAt string `json:"last_event_at"`
	Syslog      SysLog `json:"syslog"`
}
