package receiver

type PointerReceiver int

func (p *PointerReceiver) Call() {}

type ValueReceiver int

func (p ValueReceiver) Call() {}
