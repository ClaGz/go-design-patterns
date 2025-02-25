package adapter

const (
	TypeWindows = "windows"
	TypeMac     = "mac"
)

func Run(typeMachine string) {
	var adapter Computer

	switch typeMachine {
	case TypeWindows:
		adapter = &WindowsAdapter{
			windowsMachine: new(Windows),
		}
	case TypeMac:
		adapter = &MacAdapter{
			macMachine: new(Mac),
		}
	default:
		adapter = &WindowsAdapter{
			new(Windows),
		}
	}

	adapter.InsertIntoLightningPort()
}
