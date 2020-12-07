package accs

/*
Configuration Here we define the very technical settings that possibly never change and define the server “identity”.
{
	"udpPort": 9201,
	"tcpPort": 9201,
	"maxConnections": 85,
	"lanDiscovery": 1,
	"registerToLobby": 1,
	"configVersion": 1
}
The most important thing to know is that both ports must be unique on the system, the firewall allows connections and the ports are accessible from the internet.
*/
type Configuration struct {
	TCPPort         int `json:"tcpPort"`         //ACC clients will use this port to establish a connection to the server
	UDPPort         int `json:"udpPort"`         //Connected clients will use this Port to stream the car positions and is used for the ping test. In case you never see your server getting a ping value, this indicates that the udpPort is not accessible
	RegisterToLobby int `json:"registerToLobby"` //When 0, this server won’t register tothe backend. Is useful for LAN sessions. If 0, the server is declared “Private Multiplayer”.
	MaxConnections  int `json:"maxConnections"`  //Replaces “maxClients”. The maximum amount of connections a server will accept at a time. If you own the hardware server, you can just set any highnumber you want. If you rented a 16 or 24 slot server, your Hosting Provider probably has set this here and doesn’t give you write-access to this configuration file.
	LanDiscovery    int `json:"lanDiscovery"`    //Defines if the server will listen to LAN discovery requests. Can be turned off for dedicated servers.III.2.2settings.jsonThe setting defines your personal server settings, which may be changed from time to time, but also define the server.{"serverName": "Kunos Test Server #03","adminPassword": " adminPw123","carGroup": "GT4","trackMedalsRequirement": 3,"safetyRatingRequirement": 49,"racecraftRatingRequirement": -1,"password": "accessPw123","spectatorPassword": "spectPw432","maxCarSlots": 30,"dumpLeaderboards": 0,"isRaceLocked": 1,"randomizeTrackWhenEmpty": 0,"centralEntryListPath": "","allowAutoDQ": 1,"shortFormationLap": 0,"dumpEntryList": 0,"formationLapType": 3}PropertyRemarksserverNameThe server name displayed in the ACC UI pagesadminPasswordPassword to elevate via “Server admin commands”carGroupDefines the car group for this server. Possible values are"FreeForAll", "GT3", "GT4", "Cup", "ST"where “FreeForAll” will allow any driver to join with any car (that he defined as Primary Car).GT3, GT4, Cup, ST will restrict this server to GT3, GT4, Porsche Cup oder Lamborghini Supertrofeo.trackMedalsRequirementDefines the amount of track medals that a user has to have for the given track (values 0, 1, 2, 3)safetyRatingRequirementDefines the Safety Rating (SA) that a user must have to join this server (values -1, 0, .... 99)racecraftRatingRequirementDefines the Safety Rating (RC) that a user must have to join
}

/*
Settings The setting defines your personal server settings, which may be changed from time to time, but also define the server.
{
	"serverName": "Kunos Test Server #03",
	"adminPassword": " adminPw123",
	"carGroup": "GT4",
	"trackMedalsRequirement": 3,
	"safetyRatingRequirement": 49,
	"racecraftRatingRequirement": -1,
	"password": "accessPw123",
	"spectatorPassword": "spectPw432",
	"maxCarSlots": 30,
	"dumpLeaderboards": 0,
	"isRaceLocked": 1,
	"randomizeTrackWhenEmpty": 0,
	"centralEntryListPath": "",
	"allowAutoDQ": 1,
	"shortFormationLap": 0,
	"dumpEntryList": 0,
	"formationLapType": 3
}
*/
type Settings struct {
	ServerName                 string `json:"serverName"`                 //The server name displayed in the ACC UI pages
	AdminPassword              string `json:"adminPassword"`              //Password to elevate via “Server admin commands”
	CarGroup                   string `json:"carGroup"`                   //Defines the car group for this server. Possible values are"FreeForAll", "GT3", "GT4", "Cup", "ST"where “FreeForAll” will allow any driver to join with any car (that he defined as Primary Car).GT3, GT4, Cup, ST will restrict this server to GT3, GT4, Porsche Cup oder Lamborghini Supertrofeo.
	TrackMedalsRequirement     int    `json:"trackMedalsRequirement"`     //Defines the amount of track medals that a user has to have for the given track (values 0, 1, 2, 3)
	SafetyRatingRequirement    int    `json:"safetyRatingRequirement"`    //Defines the Safety Rating (SA) that a user must have to join this server (values -1, 0, .... 99)
	RacecraftRatingRequirement int    `json:"racecraftRatingRequirement"` //Defines the Safety Rating (RC) that a user must have to join this server (values -1, 0, .... 99)
	Password                   string `json:"password"`                   //Password required to enter this server. If a password is set, the server is declared “Private Multiplayer”.
	SpectatorPassword          string `json:"spectatorPassword"`          //Password to enter the server as spectator. Must be different to “password” if both is set.
	MaxCarSlots                int    `json:"maxCarSlots"`                //Replaces “maxClientsOverride” and “spectatorSlots”. Defines the amount of car slots the server can occupy; this value is overridden if the pit count of the track is lower, or with 30 for public MP. The gap between maxCarSlots and maxConnections defines howmany spectators or other irregular connections (ie entry list entries) can be on the server.
	DumpLeaderboards           int    `json:"dumpLeaderboards"`           //If set to 1, any session will write down the result leaderboard in a “results” folder (must be manually created). See ”Session results”
	IsRaceLocked               int    `json:"isRaceLocked"`               //If set to 0, the server will allow joining during a race session. Is not useful in “Public Multiplayer”, as the user-server matching will ignore ongoing race sessions.
	RandomizeTrackWhenEmpty    int    `json:"randomizeTrackWhenEmpty"`    //If set to 1, the server will change to a random track when the last drivers leaves (which causes a reset to FP1). The “track” property will only define the default state for the first session.
	CentralEntryListPath       string `json:"centralEntryListPath"`       //Can override the default entryList path “cfg/entrylist.json”, so multiple ACC servers on the machine can use the same entrylist (and custom car files). Set a full path like “C:/customEntryListSeriesA/”, where the entrylist is stored.Attention: The path seperators have to be slashes (/), backslashes (\) will not work.
	AllowAutoDQ                int    `json:"allowAutoDQ"`                //If set to 0, the server won’t automatically disqualify drivers, and instead hand out Stop&Go (30s) penalties. This way a server admin / race director has 3 laps time to review the incident, and either use /dq or /clear based on his judgement.
	ShortFormationLap          int    `json:"shortFormationLap"`          //Toggles the short and long formation lap. Long formation is only usable on private servers.
	DumpEntryList              int    `json:"dumpEntryList"`              //Will save an entry list at the end of any Qualifying session. This can be a quick way to collect a starting point to build an entry list, and is a way to save the defaultGridPositions which can be used to run a race without Qualifying session and predefined grid. Also see the corresponding admin command.
	FormationLapType           int    `json:"formationLapType"`           //Toggles the formation lap type that is permanently used on this server:3 –default formation lap with position control and UI0 –old limiter lap1 –free (replaces /manual start), only usable for private servers
}

//Event : Defines the race weekend the server runs
type Event struct {
	Track                     string  `json:"track"`
	PreRaceWaitingTimeSeconds int     `json:"preRaceWaitingTimeSeconds"`
	SessionOverTimeSeconds    int     `json:"sessionOverTimeSeconds"`
	AmbientTemp               int     `json:"ambientTemp"`
	CloudLevel                float64 `json:"cloudLevel"` //0.0, 0.1, .... 1.0
	Rain                      float64 `json:"rain"`       //0.0, 0.1, .... 1.0
	WeatherRandomness         int     `json:"weatherRandomness"`
	PostQualySeconds          int     `json:"postQualySeconds"`
	PostRaceSeconds           int     `json:"postRaceSeconds"`
	Sessions                  []struct {
		HourOfDay              int    `json:"hourOfDay"`              //Session starting hour of the day (values 0 -23)
		DayOfWeekend           int    `json:"dayOfWeekend"`           //Race day (1 = Friday, 2 = Saturday, 3 = Sunday) –relevant to the grip and weather simulation.
		TimeMultiplier         int    `json:"timeMultiplier"`         //Rate at which the session time advances in realtime. Values 0, 1, ... 24
		SessionType            string `json:"sessionType"`            //Race session type: P, Q, R for (P)ractice, (Q)ualy, (R)ace
		SessionDurationMinutes int    `json:"sessionDurationMinutes"` //Session duration in minutes
	} `json:"sessions"` //A list of session objects
	MetaData                      string `json:"metaData"` //A user defined string that will be transferred to the result outputs.
	SimracerWeatherConditions     int    `json:"simracerWeatherConditions"`
	IsFixedConditionQualification int    `json:"isFixedConditionQualification"`
}

//EventRules :Defines the pitstop rules. Public MP servers will ignore this json file and use default values.
type EventRules struct {
	QualifyStandingType                  int  `json:"qualifyStandingType"`                  //1 = fastest lap, 2 = average lap (running Endurance mode for multiple Q sessions) . Use 1, averaging Qualy is not yet officially supported.
	PitWindowLengthSec                   int  `json:"pitWindowLengthSec"`                   //Defines a pit window at the middle of the race. Obviously covers the Sprint series format. -1 will disable the pit window. Use this combined with a mandatoryPitstopCount = 1.
	DriverStintTimeSec                   int  `json:"driverStintTimeSec"`                   //Defines the maximum time a driver can stay out without getting a penalty. Can be used to balance fuel efficient cars in endurance races. The stint time resets in the pitlane, no real stop is required. -1 will disable the stint times. driverStintTimeSec and maxTotalDrivingTime are interdependent features, make sure both are set or off.
	MandatoryPitstopCount                int  `json:"mandatoryPitstopCount"`                //Defines the basic mandatory pit stops. If the value is greater zero, any car that did not execute the mandatory pitstops will be disqualified at the end of the race. The necessary actions can be further configured using the “isMandatoryPitstopXYRequired” properties. A value of zero disables the feature.
	MaxTotalDrivingTime                  int  `json:"maxTotalDrivingTime"`                  // Restricts the maximum driving time for a single driver. Is only useful for driver swap situations and allows to enforce a minimum driving timefor each driver (IRL this is used to make sure mixed teams like Pro/Am have a fair distributions of the slower drivers). -1 disables the feature. driverStintTimeSec and maxTotalDrivingTime are interdependent features, make sure both are set or off.Will set the maximum driving time for the team size defined by “maxDriversCount”, always make sure both are set.
	MaxDriversCount                      int  `json:"maxDriversCount"`                      //In driver swap situations, set this to the maximum number of drivers on a car. When an entry has fewer drivers than maxDriversCount, maxTotalDrivingTime is automatically compensated so that those "smaller" entries are also able to complete the raceExample: 3H race length, 65 minutes driverStintTimeSec and 65 minutes maxTotalDrivingTime will result in 65 minutes of maxTotalDrivingTime for entries of 3 and 105 (!) minutes for entries of 2.
	IsRefuellingAllowedInRace            bool `json:"isRefuellingAllowedInRace"`            //Defines if refuelling is allowed during the race pitstops.
	IsRefuellingTimeFixed                bool `json:"isRefuellingTimeFixed"`                //If set to true, any refuelling will take the same amount of time. If turned off, refuelling will consume time linear to the amount refuelled. Very useful setting to balance fuel efficient cars, especially if combinedwith other features.
	IsMandatoryPitstopRefuellingRequired bool `json:"isMandatoryPitstopRefuellingRequired"` //Defines if a mandatory pitstop requires refuelling.
	IsMandatoryPitstopTyreChangeRequired bool `json:"isMandatoryPitstopTyreChangeRequired"` //Defines if a mandatory pitstop requires changing tyres.
	IsMandatoryPitstopSwapDriverRequired bool `json:"isMandatoryPitstopSwapDriverRequired"` //Defines if a mandatory pitstop requires a driver swap. Will only be effective for cars in driver swap situations; even in a mixed field this will be skipped for cars with a team size of 1 driver.
	TyreSetCount                         int  `json:"tyreSetCount"`                         //Experimental/not supported: Can be used to reduce the amount of tyre sets any car entry has for the entire weekend. Please note that it is necessary to force cars to remain in the server, or drastically reduced tyre sets will be ineffective, as rejoining will resetthe tyre sets.
}

//AssistRules : Can be used to turn off certain assists for any car connected to this server. Beware: disabling assists will effectively remove the effect, but there is no special handling how the assists look like in the menu. Without instructions, users will be surprised and confused –up to a point where they become a risk for other drivers. Whenever you think about disabling something, please be sure this is really necessary and a risk in terms of fairness. It is out of question that the (quite strong) driving aids “Stability Control” and “Autosteer” may be candidates for league racing, but just turning off the ideal line will not improve anything for anyone (except that the one driver using it may become less safe and ruins the race of others). Even innocent elements like auto-engine start and pit limiter may just force users to re-map their wheels, and for example lose the ability to use their indicators in lapping traffic –again nobody is winning in this scenario
type AssistRules struct {
	StabilityControlLevelMax int `json:"stabilityControlLevelMax"` //Set’s the maximum % of SC that can be used. In case a client has a higher SC set than allowed by the server, he will only run what is allowed (25% in this example). Obviously setting this property to 0 removes all SC, including mouse and keyboard users.The Stability Control is an artificial driving aid that allows the car to act out of the physics boundaries, and highly recommended to overcome input methods like Keyboards, Gamepads and Mouse steering. However, there is a built-in effect that makes the SC performance inferior, so in theory using (and relying) on SC is already more than enough penalty, and the way to improve performance is to practice driving without.Default: 100
	DisableAutosteer         int `json:"disableAutosteer"`         //Disables the steering aid that is only available for gamepad controllers. Unlike SC, this works inside the physics and does not allow unrealistic driving behaviour –except that this is a very strong aid with superhuman feeling for grip and high reaction speed. There is a built-in penalty that should balance the driving performance in most cases, and give an incentive to learn not to use the driving aid.Default: 0
	DisableAutoLights        int `json:"disableAutoLights"`
	DisableAutoWiper         int `json:"disableAutoWiper"`
	DisableAutoEngineStart   int `json:"disableAutoEngineStart"`
	DisableAutoPitLimiter    int `json:"disableAutoPitLimiter"`
	DisableAutoGear          int `json:"disableAutoGear"`
	DisableAutoClutch        int `json:"disableAutoClutch"`
	DisableIdealLine         int `json:"disableIdealLine"`
}
