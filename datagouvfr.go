package godatagouvfr

import (
	"errors"
)

func GetDataGouvFrEntityByRna(rna string) (AssociationEntity, error) {
	var info SiretResponseAssociation
	err := JsonRequest("https://entreprise.data.gouv.fr/api/rna/v1/id/"+rna, "GET", &info)
	if err != nil {
		return AssociationEntity{}, err
	}
	if info.SiretResp.Id == 0 {
		return AssociationEntity{}, errors.New("association not found")
	}
	info.SiretResp.Rna = rna
	return AssociationEntity(info.SiretResp), nil
}

func GetDataGouvFrEntityBySiret(siret string) (AssociationEntity, error) {
	var info RnaResponseAssociation
	err := JsonRequest("https://entreprise.data.gouv.fr/api/rna/v1/siret/"+siret, "GET", &info)
	if err != nil {
		return AssociationEntity{}, err
	}
	if info.RnaResp.Siret == "" {
		return AssociationEntity{}, errors.New("association not found")
	}
	return AssociationEntity(info.RnaResp), nil
}

func GetDataGouvFrEntityPageByFullText(text string) (AssociationEntityPage, error) {
	var info TextResponseAssociation
	err := JsonRequest("https://entreprise.data.gouv.fr/api/rna/v1/full_text/"+text, "GET", &info)
	if err != nil {
		return AssociationEntityPage{}, err
	}
	entityPage := AssociationEntityPage{
		Count:  info.Count,
		Cursor: info.Cursor,
		Items:  make([]AssociationEntity, len(info.Items)),
	}
	for i, s := range info.Items {
		entityPage.Items[i] = AssociationEntity(s)
	}
	return entityPage, nil
}
