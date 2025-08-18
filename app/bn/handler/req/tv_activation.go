package req

type TVActivationRequest struct {
	Activate   []ActivationRequest `json:"activate"`
	Deactivate []ActivationRequest `json:"deactivate"`
}
