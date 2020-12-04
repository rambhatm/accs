package accs

//Configuration : contains connection related parameters like ports
type Configuration struct {
	TCPPort         int `json:"tcpPort"`
	UDPPort         int `json:"udpPort"`
	RegisterToLobby int `json:"registerToLobby"`
	MaxConnections  int `json:"maxConnections"`
	LanDiscovery    int `json:"lanDiscovery"`
}

//Settings : list of parameters for generic server settings
type Settings struct {
	ServerName                 string `json:"serverName"`
	AdminPassword              string `json:"adminPassword"`
	CarGroup                   string `json:"carGroup"`                   //"FreeForAll", "GT3", "GT4", "Cup", "ST"
	TrackMedalsRequirement     int    `json:"trackMedalsRequirement"`     //0, 1, 2, 3
	SafetyRatingRequirement    int    `json:"safetyRatingRequirement"`    //-1, 0, …. 99
	RacecraftRatingRequirement int    `json:"racecraftRatingRequirement"` //-1, 0, …. 99
	Password                   string `json:"password"`
	SpectatorPassword          string `json:"spectatorPassword"`
	MaxCarSlots                int    `json:"maxCarSlots"`
	DumpLeaderboards           int    `json:"dumpLeaderboards"`
	IsRaceLocked               int    `json:"isRaceLocked"`
	RandomizeTrackWhenEmpty    int    `json:"randomizeTrackWhenEmpty"`
	CentralEntryListPath       string `json:"centralEntryListPath"`
	AllowAutoDQ                int    `json:"allowAutoDQ"`
	ShortFormationLap          int    `json:"shortFormationLap"`
	DumpEntryList              int    `json:"dumpEntryList"`
	FormationLapType           int    `json:"formationLapType"` //3 –default formation lap with position control and UI0 –old limiter lap1 –free (replaces /manual start), only usable for private servers
}
