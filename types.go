package godatagouvfr

type NpoPage struct {
	Count  int    `json:"count"`
	Cursor string `json:"cursor"`
	Items  []Npo  `json:"items"`
}

type Npo struct {
	Id        int    `json:"id"`
	Rna       string `json:"rna"`
	Name      string `json:"name"`
	ShortName string `json:"short_Name"`
	ZipCode   string `json:"zipCode"`
	Address   string `json:"address"`
	Country   string `json:"country"`
	Siret     string `json:"siret"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type RnaResponseAssociation struct {
	RnaResp Response `json:"association"`
}

type SiretResponseAssociation struct {
	SiretResp Response `json:"association"`
}

type TextResponseAssociation struct {
	Count  int        `json:"total_results"`
	Cursor string     `json:"cursor"`
	Items  []Response `json:"association"`
}

type Response struct {
	Id        int    `json:"id"`
	Rna       string `json:"id_association"`
	Name      string `json:"titre"`
	ShortName string `json:"titre_court"`
	ZipCode   string `json:"adresse_code_postal"`
	Address   string `json:"adresse_siege"`
	Country   string `json:"adresse_gestion_pays"`
	Siret     string `json:"siret"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AssociationEntity struct {
	Id        int    `json:"id"`
	Rna       string `json:"rna"`
	Name      string `json:"name"`
	ShortName string `json:"short_Name"`
	ZipCode   string `json:"zipCode"`
	Address   string `json:"address"`
	Country   string `json:"country"`
	Siret     string `json:"siret"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AssociationEntityPage struct {
	Count  int                 `json:"count"`
	Cursor string              `json:"cursor"`
	Items  []AssociationEntity `json:"items"`
}
