package dto

type ClientRequest struct {
	ClientName   string ` json:"clientName"`
	ClientActive bool   ` json:"active"`
	ClientTel    string ` json:"telefone"`
}
