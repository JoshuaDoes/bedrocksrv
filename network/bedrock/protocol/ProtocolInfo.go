package protocol

//Server protocol constants
const (
	//CurrentProtocol contains the protocol version to tell the client
	CurrentProtocol = 291

	//MinecraftVersion contains the version that will be reported by the server
	//This is usually the earliest currently supported version
	MinecraftVersion = "v1.7.0"

	//MinecraftVersionNetwork contains the version number sent to clients in ping responses
	MinecraftVersionNetwork = "1.7.0"
)

//Packet IDs for communicating with clients
const (
	//OpLogin contains the Login packet ID
	OpLogin = 0x01

	//OpPlayStatus contains the Play Status packet ID
	OpPlayStatus = 0x02

	//OpServerToClientHandshake contains the Server To Client Handshake packet ID
	OpServerToClientHandshake = 0x03

	//OpClientToServerHandshake contains the Client To Server Handshake packet ID
	OpClientToServerHandshake = 0x04

	//OpDisconnect contains the Disconnect packet ID
	OpDisconnect = 0x05

	//OpResourcePacksInfo contains the Resource Packs Info packet ID
	OpResourcePacksInfo = 0x06

	//OpResourcePackStack contains the Resource Pack Stack packet ID
	OpResourcePackStack = 0x07

	//OpResourcePackClientResponse contains the Resource Pack Client Response packet ID
	OpResourcePackClientResponse = 0x08

	//OpText contains the Text packet ID
	OpText = 0x09

	//OpSetTime contains the Set Time packet ID
	OpSetTime = 0x0A

	//OpStartGame contains the Start Game packet ID
	OpStartGame = 0x0B

	//OpAddPlayer contains the Add Player packet ID
	OpAddPlayer = 0x0C

	//OpAddEntity contains the Add Entity packet ID
	OpAddEntity = 0x0D

	//OpRemoveEntity contains the Remove Entity packet ID
	OpRemoveEntity = 0x0E

	//OpAddItemEntity contains the Add Item Entity packet ID
	OpAddItemEntity = 0x0F

	//OpAddHangingEntity contains the Add Hanging Entity packet ID
	OpAddHangingEntity = 0x10

	//OpTakeItemEntity contains the Take Item Entity packet ID
	OpTakeItemEntity = 0x11

	//OpMoveEntityAbsolute contains the Move Entity Absolute packet ID
	OpMoveEntityAbsolute = 0x12

	//OpMovePlayer contains the Move Player packet ID
	OpMovePlayer = 0x13

	//OpRiderJump contains the Rider Jump packet ID
	OpRiderJump = 0x14

	//OpUpdateBlock contains the Update Block packet ID
	OpUpdateBlock = 0x15

	//OpAddPainting contains the Add Painting packet ID
	OpAddPainting = 0x16

	//OpExplode contains the Explode packet ID
	OpExplode = 0x17

	//OpLevelSoundEvent contains the Level Sound Event packet ID
	OpLevelSoundEvent = 0x18

	//OpLevelEvent contains the Level Event packet ID
	OpLevelEvent = 0x19

	//OpBlockEvent contains the Block Event packet ID
	OpBlockEvent = 0x1A

	//OpEntityEvent contains the Entity Event packet ID
	OpEntityEvent = 0x1B

	//OpMobEffect contains the Mob Effect packet ID
	OpMobEffect = 0x1C

	//OpUpdateAttributes contains the Update Attributes packet ID
	OpUpdateAttributes = 0x1D

	//OpInventoryTransaction contains the Inventory Transaction packet ID
	OpInventoryTransaction = 0x1E

	//OpMobEquipment contains the Mob Equipment packet ID
	OpMobEquipment = 0x1F

	//OpMobArmorEquipment contains the Mob Armor Equipment packet ID
	OpMobArmorEquipment = 0x20

	//OpInteract contains the Interact packet ID
	OpInteract = 0x21

	//OpBlockPickRequest contains the Block Pick Request packet ID
	OpBlockPickRequest = 0x22

	//OpEntityPickRequest contains the Entity Pick Request packet ID
	OpEntityPickRequest = 0x23

	//OpPlayerAction contains the Player Action packet ID
	OpPlayerAction = 0x24

	//OpEntityFall contains the Entity Fall packet ID
	OpEntityFall = 0x25

	//OpHurtArmor contains the Hurt Armor packet ID
	OpHurtArmor = 0x26

	//OpSetEntityData contains the Set Entity Data packet ID
	OpSetEntityData = 0x27

	//OpSetEntityMotion contains the Set Entity Motion packet ID
	OpSetEntityMotion = 0x28

	//OpSetEntityLink contains the Set Entity Link packet ID
	OpSetEntityLink = 0x29

	//OpSetHealth contains the Set Health packet ID
	OpSetHealth = 0x2A

	//OpSetSpawnPosition contains the Set Spawn Position packet ID
	OpSetSpawnPosition = 0x2B

	//OpAnimate contains the Animate packet ID
	OpAnimate = 0x2C

	//OpRespawn contains the Respawn packet ID
	OpRespawn = 0x2D

	//OpContainerOpen contains the Container Open packet ID
	OpContainerOpen = 0x2E

	//OpContainerClose contains the Container Close packet ID
	OpContainerClose = 0x2F

	//OpPlayerHotbar contains the Player Hotbar packet ID
	OpPlayerHotbar = 0x30

	//OpInventoryContent contains the Inventory Content packet ID
	OpInventoryContent = 0x31

	//OpInventorySlot contains the Inventory Slot packet ID
	OpInventorySlot = 0x32

	//OpContainerSetData contains the Container Set Data packet ID
	OpContainerSetData = 0x33

	//OpCraftingData contains the Crafting Data packet ID
	OpCraftingData = 0x34

	//OpCraftingEvent contains the Crafting Event packet ID
	OpCraftingEvent = 0x35

	//OpGUIDataPickItem contains the GUI Data Pick Item packet ID
	OpGUIDataPickItem = 0x36

	//OpAdventureSettings contains the Adventure Settings packet ID
	OpAdventureSettings = 0x37

	//OpBlockEntityData contains the Block Entity Data packet ID
	OpBlockEntityData = 0x38

	//OpPlayerInput contains the Player Input packet ID
	OpPlayerInput = 0x39

	//OpFullChunkData contains the Full Chunk Data packet ID
	OpFullChunkData = 0x3A

	//OpSetCommandsEnabled contains the Set Commands Enabled packet ID
	OpSetCommandsEnabled = 0x3B

	//OpSetDifficulty contains the Set Difficulty packet ID
	OpSetDifficulty = 0x3C

	//OpChangeDimension contains the Change Dimension packet ID
	OpChangeDimension = 0x3D

	//OpSetPlayerGameType contains the Set Player Game Type packet ID
	OpSetPlayerGameType = 0x3E

	//OpPlayerList contains the Player List packet ID
	OpPlayerList = 0x3F

	//OpSimpleEvent contains the Simple Event packet ID
	OpSimpleEvent = 0x40

	//OpEvent contains the Event packet ID
	OpEvent = 0x41

	//OpSpawnExperienceOrb contains the Spawn Experience Orb packet ID
	OpSpawnExperienceOrb = 0x42

	//OpClientboundMapItemData contains the Clientbound Map Item Data packet ID
	OpClientboundMapItemData = 0x43

	//OpMapInfoRequest contains the Map Info Request packet ID
	OpMapInfoRequest = 0x44

	//OpRequestChunkRadius contains the Request Chunk Radius packet ID
	OpRequestChunkRadius = 0x45

	//OpChunkRadiusUpdated contains the Chunk Radius Updated packet ID
	OpChunkRadiusUpdated = 0x46

	//OpItemFrameDropItem contains the Item Frame Drop Item packet ID
	OpItemFrameDropItem = 0x47

	//OpGameRulesChanged contains the Game Rules Changed packet ID
	OpGameRulesChanged = 0x48

	//OpCamera contains the Camera packet ID
	OpCamera = 0x49

	//OpBossEvent contains the Boss Event packet ID
	OpBossEvent = 0x4A

	//OpShowCredits contains the Show Credits packet ID
	OpShowCredits = 0x4B

	//OpAvailableCommands contains the Available Commands packet ID
	OpAvailableCommands = 0x4C

	//OpCommandRequest contains the Command Request packet ID
	OpCommandRequest = 0x4D

	//OpCommandBlockUpdate contains the Command Block Update packet ID
	OpCommandBlockUpdate = 0x4E

	//OpCommandOutput contains the Command Output packet ID
	OpCommandOutput = 0x4F

	//OpUpdateTrade contains the Update Trade packet ID
	OpUpdateTrade = 0x50

	//OpUpdateEquip contains the Update Equip packet ID
	OpUpdateEquip = 0x51

	//OpResourcePackDataInfo contains the Resource Pack Data Info packet ID
	OpResourcePackDataInfo = 0x52

	//OpResourcePackChunkData contains the Resource Pack Chunk Data packet ID
	OpResourcePackChunkData = 0x53

	//OpResourcePackChunkRequest contains the Resource Pack Chunk Request packet ID
	OpResourcePackChunkRequest = 0x54

	//OpTransfer contains the Transfer packet ID
	OpTransfer = 0x55

	//OpPlaySound contains the Play Sound packet ID
	OpPlaySound = 0x56

	//OpStopSound contains the Stop Sound packet ID
	OpStopSound = 0x57

	//OpSetTitle contains the Set Title packet ID
	OpSetTitle = 0x58

	//OpAddBehaviorTree contains the Add Behavior Tree packet ID
	OpAddBehaviorTree = 0x59

	//OpStructureBlockUpdate contains the Structure Block Update packet ID
	OpStructureBlockUpdate = 0x5A

	//OpShowStoreOffer contains the Show Store Offer packet ID
	OpShowStoreOffer = 0x5B

	//OpPurchaseReceipt contains the Purchase Receipt packet ID
	OpPurchaseReceipt = 0x5C

	//OpPlayerSkin contains the Player Skin packet ID
	OpPlayerSkin = 0x5D

	//OpSubClientLogin contains the Sub Client Login packet ID
	OpSubClientLogin = 0x5E

	//OpWSConnect contains the W S Connect packet ID
	OpWSConnect = 0x5F

	//OpSetLastHurtBy contains the Set Last Hurt By packet ID
	OpSetLastHurtBy = 0x60

	//OpBookEdit contains the Book Edit packet ID
	OpBookEdit = 0x61

	//OpNPCRequest contains the NPC Request packet ID
	OpNPCRequest = 0x62

	//OpPhotoTransfer contains the Photo Transfer packet ID
	OpPhotoTransfer = 0x63

	//OpModalFormRequest contains the Modal Form Request packet ID
	OpModalFormRequest = 0x64

	//OpModalFormResponse contains the Modal Form Response packet ID
	OpModalFormResponse = 0x65

	//OpServerSettingsRequest contains the Server Settings Request packet ID
	OpServerSettingsRequest = 0x66

	//OpServerSettingsResponse contains the Server Settings Response packet ID
	OpServerSettingsResponse = 0x67

	//OpShowProfile contains the Show Profile packet ID
	OpShowProfile = 0x68

	//OpSetDefaultGameType contains the Set Default Game Type packet ID
	OpSetDefaultGameType = 0x69

	//OpRemoveObjective contains the Remove Objective packet ID
	OpRemoveObjective = 0x6A

	//OpSetDisplayObjective contains the Set Display Objective packet ID
	OpSetDisplayObjective = 0x6B

	//OpSetScore contains the Set Score packet ID
	OpSetScore = 0x6C

	//OpLabTable contains the Lab Table packet ID
	OpLabTable = 0x6D

	//OpUpdateBlockSynced contains the Update Block Synced packet ID
	OpUpdateBlockSynced = 0x6E

	//OpMoveEntityDelta contains the Move Entity Delta packet ID
	OpMoveEntityDelta = 0x6F

	//OpSetScoreboardIdentity contains the Set Scoreboard Identity packet ID
	OpSetScoreboardIdentity = 0x70

	//OpSetLocalPlayerAsInitialized contains the Set Local Player As Initialized packet ID
	OpSetLocalPlayerAsInitialized = 0x71

	//OpUpdateSoftEnum contains the Update Soft Enum packet ID
	OpUpdateSoftEnum = 0x72

	//OpNetworkStackLatency contains the Network Stack Latency packet ID
	OpNetworkStackLatency = 0x73

	//OpScriptCustomEvent contains the Script Custom Event packet ID
	OpScriptCustomEvent = 0x75
)
