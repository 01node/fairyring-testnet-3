package types

// ValidateBasic is used for validating the packet
func (p CurrentKeysPacketData) ValidateBasic() error {
	return nil
}

// GetBytes is a helper for serialising
func (p CurrentKeysPacketData) GetBytes() ([]byte, error) {
	var modulePacket ConditionalencPacketData

	modulePacket.Packet = &ConditionalencPacketData_CurrentKeysPacket{&p}

	return modulePacket.Marshal()
}
