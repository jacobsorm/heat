package heat

type DeviceResponse struct {
	ID                string            `json:"id"`
	ConnectionState   string            `json:"connectionState"`
	Firmware          Firmware          `json:"firmware"`
	Product           Product           `json:"product"`
	AvailableFeatures AvailableFeatures `json:"availableFeatures"`
}

type Firmware struct {
	CurrentFwVersion string `json:"currentFwVersion"`
	DesiredFwVersion string `json:"desiredFwVersion"`
}

type Product struct {
	SerialNumber string `json:"serialNumber"`
	Name         string `json:"name"`
}

type AvailableFeatures struct {
	Settings            bool `json:"settings"`
	Reboot              bool `json:"reboot"`
	ForceSync           bool `json:"forcesync"`
	ForceUpdate         bool `json:"forceUpdate"`
	RequestUpdate       bool `json:"requestUpdate"`
	ResetAlarm          bool `json:"resetAlarm"`
	TriggerEvent        bool `json:"triggerEvent"`
	GetMenu             bool `json:"getMenu"`
	GetMenuChain        bool `json:"getMenuChain"`
	GetGuideQuestion    bool `json:"getGuideQuestion"`
	SendHaystack        bool `json:"sendHaystack"`
	SetSmartMode        bool `json:"setSmartMode"`
	SetAidMode          bool `json:"setAidMode"`
	GetZones            bool `json:"getZones"`
	ProcessIntent       bool `json:"processIntent"`
	BoostHotWater       bool `json:"boostHotWater"`
	BoostVentilation    bool `json:"boostVentilation"`
	GetScheduleConfig   bool `json:"getScheduleConfig"`
	GetScheduleModes    bool `json:"getScheduleModes"`
	GetScheduleWeekly   bool `json:"getScheduleWeekly"`
	GetScheduleVacation bool `json:"getScheduleVacation"`
	SetScheduleModes    bool `json:"setScheduleModes"`
	SetScheduleWeekly   bool `json:"setScheduleWeekly"`
	SetScheduleOverride bool `json:"setScheduleOverride"`
	SetScheduleVacation bool `json:"setScheduleVacation"`
	SetVentilationMode  bool `json:"setVentilationMode"`
}

type Signal struct {
	Name string
	Type SignalType
	Enum map[int]string
}

type SignalType int

const (
	SignalTypeTemperature SignalType = iota
	SignalTypeCurrent
	SignalTypeDays
	SignalTypeHours
	SignalTypeMinutes
	SignalTypeSeconds
	SignalTypePercent
	SignalTypeNone
)

type SignalId string

var (
	Signal40004OutdoorTemperature               SignalId = "40004"
	Signal40008SupplyLineBT2                    SignalId = "40008"
	Signal40012ReturnLineBT3                    SignalId = "40012"
	Signal40013HotWaterTopBT7                   SignalId = "40013"
	Signal40014HotWaterChargingBT6              SignalId = "40014"
	Signal40020EvaporatorBT16                   SignalId = "40020"
	Signal40023CompressorSensorBT18             SignalId = "40023"
	Signal40024ImmersionHeaterSensorBT19        SignalId = "40024"
	Signal40025ExhaustAirAZ30BT20               SignalId = "40025"
	Signal40026ExtractAirAZ30BT21               SignalId = "40026"
	Signal40067AverageOutdoorTempBT1            SignalId = "40067"
	Signal40079CurrentBE3                       SignalId = "40079"
	Signal40081CurrentBE2                       SignalId = "40081"
	Signal40083CurrentBE1                       SignalId = "40083"
	Signal40145OilTemperatureEP15BT29           SignalId = "40145"
	Signal40146OilTemperatureBT29               SignalId = "40146"
	Signal41929ModeSmartPriceAdaption           SignalId = "41929"
	Signal42770DesiredHumidity                  SignalId = "42770"
	Signal43009CalculatedSupplyClimateSystem1   SignalId = "43009"
	Signal43066DefrostingTime                   SignalId = "43066"
	Signal43081TimeFactorAddHeat                SignalId = "43081"
	Signal43108CurrentFanMode                   SignalId = "43108"
	Signal43109CurrentHotWaterMode              SignalId = "43109"
	Signal43161ExternalAdjustmentClimateSystem1 SignalId = "43161"
	Signal43416NumberOfCompressorStarts         SignalId = "43416"
	Signal43420OperatingTime                    SignalId = "43420"
	Signal43424OpertingTimeHotWater             SignalId = "43424"
	Signal43427StatusCompressor                 SignalId = "43427"
	Signal44896HeatingOffsetSmartPriceAdaption  SignalId = "44896"
	Signal44898PoolOffsetSmartPriceAdaption     SignalId = "44898"
	Signal44899CoolingOffsetSmartPriceAdaption  SignalId = "44899"
	Signal44908StatusSmartPriceAdaption         SignalId = "44908"
	Signal47007HeatingCurve                     SignalId = "47007"
	Signal47011Offset                           SignalId = "47011"
	Signal47015ClimateSystem                    SignalId = "47015"
	Signal47020FlowLineTempAt30C                SignalId = "47020"
	Signal47021FlowLineTempAt20C                SignalId = "47021"
	Signal47022FlowLineTempAt10C                SignalId = "47022"
	Signal47023FlowLineTempAt0C                 SignalId = "47023"
	Signal47024FlowLineTempAtMinus10C           SignalId = "47024"
	Signal47025FlowLineTempAtMinus20C           SignalId = "47025"
	Signal47026FlowLineTempAtMinus30C           SignalId = "47026"
	Signal47027OutdoorTempPoint                 SignalId = "47027"
	Signal47028ChangeInCurve                    SignalId = "47028"
	Signal47032ClimateSystem                    SignalId = "47032"
	Signal47041HotWaterAmount                   SignalId = "47041"
	Signal47050Activated                        SignalId = "47050"
	Signal47051Period                           SignalId = "47051"
	Signal47137OperatingMode                    SignalId = "47137"
	Signal47212SetMaxElectricalAdd              SignalId = "47212"
	Signal47271Speed4                           SignalId = "47271"
	Signal47272Speed3                           SignalId = "47272"
	Signal47273Speed2                           SignalId = "47273"
	Signal47274Speed1                           SignalId = "47274"
	Signal47375StopHeating                      SignalId = "47375"
	Signal47376StopAdditionalHeat               SignalId = "47376"
	Signal47377FilteringTime                    SignalId = "47377"
	Signal47394ControlRoomSensorSystem          SignalId = "47394"
	Signal47537NightCooling                     SignalId = "47537"
	Signal47538StartTempExhaustAir              SignalId = "47538"
	Signal47539MindiffOutdoorExhaust            SignalId = "47539"
	Signal48132HotWaterBoost                    SignalId = "48132"
	Signal48442Activated                        SignalId = "48442"
	Signal48443AffectRoomTemperature            SignalId = "48443"
	Signal48445AffectHotWater                   SignalId = "48445"
	Signal48566DegreeOfEffect                   SignalId = "48566"
	Signal49633DesiredHumidity                  SignalId = "49633"
	Signal49909Starttime                        SignalId = "49909"
	Signal49993IntelecAddHeat                   SignalId = "49993"
	Signal49995PumpHeatingMediumGP1             SignalId = "49995"
	Signal50004Temporarylux                     SignalId = "50004"
	Signal50005IncreasedVentilation             SignalId = "50005"
	Signal50095Status                           SignalId = "50095"
	Signal50221FanspeedExhaustAir               SignalId = "50221"
	Signal50249Textnot1                         SignalId = "50249"
	Signal50325HotwaterTemp                     SignalId = "50325"
	Signal65535Textnot2                         SignalId = "65535"
)

var Signals = map[SignalId]Signal{
	Signal40004OutdoorTemperature:               {Name: "Outdoor temperature", Type: SignalTypeTemperature},
	Signal40008SupplyLineBT2:                    {Name: "Supply line (BT2)", Type: SignalTypeTemperature},
	Signal40012ReturnLineBT3:                    {Name: "Return line (BT3)", Type: SignalTypeTemperature},
	Signal40013HotWaterTopBT7:                   {Name: "Hot water top (BT7)", Type: SignalTypeTemperature},
	Signal40014HotWaterChargingBT6:              {Name: "Hot water charging (BT6)", Type: SignalTypeTemperature},
	Signal40020EvaporatorBT16:                   {Name: "Evaporator (BT16)", Type: SignalTypeTemperature},
	Signal40023CompressorSensorBT18:             {Name: "Compressor sensor (BT18)", Type: SignalTypeTemperature},
	Signal40024ImmersionHeaterSensorBT19:        {Name: "Immersion heater sensor (BT19)", Type: SignalTypeTemperature},
	Signal40025ExhaustAirAZ30BT20:               {Name: "Exhaust air (AZ30-BT20)", Type: SignalTypeTemperature},
	Signal40026ExtractAirAZ30BT21:               {Name: "Extract air (AZ30-BT21)", Type: SignalTypeTemperature},
	Signal40067AverageOutdoorTempBT1:            {Name: "Average outdoor temp (BT1)", Type: SignalTypeTemperature},
	Signal40079CurrentBE3:                       {Name: "Current (BE3)", Type: SignalTypeCurrent},
	Signal40081CurrentBE2:                       {Name: "Current (BE2)", Type: SignalTypeCurrent},
	Signal40083CurrentBE1:                       {Name: "Current (BE1)", Type: SignalTypeCurrent},
	Signal40145OilTemperatureEP15BT29:           {Name: "Oil temperature (EP15-BT29)", Type: SignalTypeTemperature},
	Signal40146OilTemperatureBT29:               {Name: "Oil temperature (BT29)", Type: SignalTypeTemperature},
	Signal41929ModeSmartPriceAdaption:           {Name: "Mode (Smart Price Adaption)", Type: SignalTypeNone, Enum: map[int]string{0: "Unknown", 1: "Low", 2: "Normal", 3: "High"}},
	Signal42770DesiredHumidity:                  {Name: "Desired humidity", Type: SignalTypePercent},
	Signal43009CalculatedSupplyClimateSystem1:   {Name: "Calculated supply climate system 1", Type: SignalTypeTemperature},
	Signal43066DefrostingTime:                   {Name: "Defrosting time", Type: SignalTypeSeconds},
	Signal43081TimeFactorAddHeat:                {Name: "Time factor add heat", Type: SignalTypeNone},
	Signal43108CurrentFanMode:                   {Name: "Current fan mode", Type: SignalTypeNone},
	Signal43109CurrentHotWaterMode:              {Name: "Current hot water mode", Type: SignalTypeNone},
	Signal43161ExternalAdjustmentClimateSystem1: {Name: "Extern. adjustment climate system 1", Type: SignalTypeNone},
	Signal43416NumberOfCompressorStarts:         {Name: "Number of compressor starts", Type: SignalTypeNone},
	Signal43420OperatingTime:                    {Name: "Oper. time", Type: SignalTypeHours},
	Signal43424OpertingTimeHotWater:             {Name: "Oper. time hot water", Type: SignalTypeHours},
	Signal43427StatusCompressor:                 {Name: "Status compressor", Type: SignalTypeNone, Enum: map[int]string{20: "Off", 40: "Starting", 60: "Operating", 100: "Stopping"}},
	Signal44896HeatingOffsetSmartPriceAdaption:  {Name: "Heating offset (Smart Price Adaption)", Type: SignalTypeNone},
	Signal44898PoolOffsetSmartPriceAdaption:     {Name: "Pool offset (Smart Price Adaption)", Type: SignalTypeNone},
	Signal44899CoolingOffsetSmartPriceAdaption:  {Name: "Cooling offset (Smart Price Adaption)", Type: SignalTypeNone},
	Signal44908StatusSmartPriceAdaption:         {Name: "Status (Smart Price Adaption)", Type: SignalTypeNone},
	Signal47007HeatingCurve:                     {Name: "heating curve", Type: SignalTypeNone},
	Signal47011Offset:                           {Name: "Offset", Type: SignalTypeNone},
	Signal47015ClimateSystem:                    {Name: "climate system", Type: SignalTypeTemperature},
	Signal47020FlowLineTempAt30C:                {Name: "flow line temp. at 30 \u00b0C", Type: SignalTypeTemperature},
	Signal47021FlowLineTempAt20C:                {Name: "flow line temp. at 20 \u00b0C", Type: SignalTypeTemperature},
	Signal47022FlowLineTempAt10C:                {Name: "flow line temp. at 10 \u00b0C", Type: SignalTypeTemperature},
	Signal47023FlowLineTempAt0C:                 {Name: "flow line temp. at 0 \u00b0C", Type: SignalTypeTemperature},
	Signal47024FlowLineTempAtMinus10C:           {Name: "flow line temp. at -10 \u00b0C", Type: SignalTypeTemperature},
	Signal47025FlowLineTempAtMinus20C:           {Name: "flow line temp. at -20 \u00b0C", Type: SignalTypeTemperature},
	Signal47026FlowLineTempAtMinus30C:           {Name: "flow line temp. at -30 \u00b0C", Type: SignalTypeTemperature},
	Signal47027OutdoorTempPoint:                 {Name: "outdoor temp. point", Type: SignalTypeTemperature},
	Signal47028ChangeInCurve:                    {Name: "change in curve", Type: SignalTypeTemperature},
	Signal47032ClimateSystem:                    {Name: "climate system", Type: SignalTypeNone},
	Signal47041HotWaterAmount:                   {Name: "Hot water amount", Type: SignalTypeNone, Enum: map[int]string{0: "Economy", 1: "Normal", 2: "Lux"}},
	Signal47050Activated:                        {Name: "activated", Type: SignalTypeNone},
	Signal47051Period:                           {Name: "period", Type: SignalTypeDays},
	Signal47137OperatingMode:                    {Name: "op. mode", Type: SignalTypeNone},
	Signal47212SetMaxElectricalAdd:              {Name: "set max electrical add.", Type: SignalTypeNone},
	Signal47271Speed4:                           {Name: "speed 4", Type: SignalTypeHours},
	Signal47272Speed3:                           {Name: "speed 3", Type: SignalTypeHours},
	Signal47273Speed2:                           {Name: "speed 2", Type: SignalTypeHours},
	Signal47274Speed1:                           {Name: "speed 1", Type: SignalTypeHours},
	Signal47375StopHeating:                      {Name: "stop heating", Type: SignalTypeTemperature},
	Signal47376StopAdditionalHeat:               {Name: "stop additional heat", Type: SignalTypeTemperature},
	Signal47377FilteringTime:                    {Name: "filtering time", Type: SignalTypeHours},
	Signal47394ControlRoomSensorSystem:          {Name: "control room sensor syst", Type: SignalTypeNone},
	Signal47537NightCooling:                     {Name: "night cooling", Type: SignalTypeNone},
	Signal47538StartTempExhaustAir:              {Name: "start temp. exhaust air", Type: SignalTypeTemperature},
	Signal47539MindiffOutdoorExhaust:            {Name: "min diff. outdoor-exhaust", Type: SignalTypeTemperature},
	Signal48132HotWaterBoost:                    {Name: "Hot water boost", Type: SignalTypeNone},
	Signal48442Activated:                        {Name: "activated", Type: SignalTypeNone},
	Signal48443AffectRoomTemperature:            {Name: "affect room temperature", Type: SignalTypeNone},
	Signal48445AffectHotWater:                   {Name: "affect hot water", Type: SignalTypeNone},
	Signal48566DegreeOfEffect:                   {Name: "- degree of effect", Type: SignalTypeNone},
	Signal49633DesiredHumidity:                  {Name: "Desired humidity", Type: SignalTypePercent},
	Signal49909Starttime:                        {Name: "start time", Type: SignalTypeNone},
	Signal49993IntelecAddHeat:                   {Name: "Int elec add heat", Type: SignalTypeNone, Enum: map[int]string{0: "Alarm", 1: "Alarm", 2: "Active", 3: "Off", 4: "Blocked", 5: "Off", 6: "Active"}},
	Signal49995PumpHeatingMediumGP1:             {Name: "Pump: Heating medium (GP1)", Type: SignalTypeNone},
	Signal50004Temporarylux:                     {Name: "Temporary lux", Type: SignalTypeNone},
	Signal50005IncreasedVentilation:             {Name: "Increased ventilation", Type: SignalTypeNone},
	Signal50095Status:                           {Name: "status:", Type: SignalTypeMinutes}, // ??
	Signal50221FanspeedExhaustAir:               {Name: "fan speed exhaust air", Type: SignalTypePercent},
	Signal50249Textnot1:                         {Name: "Text not found: id[], fw[nibe-e2], lang[en-US]", Type: SignalTypeNone},
	Signal50325HotwaterTemp:                     {Name: "Hot water temp.", Type: SignalTypeTemperature},
	Signal65535Textnot2:                         {Name: "Text not found: id[], fw[nibe-e2], lang[en-US]", Type: SignalTypeNone},
}
