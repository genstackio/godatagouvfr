package godatagouvfr

//goland:noinspection GoUnusedExportedFunction
func GetNpoByRna(rna string) (Npo, error) {
	entity, err := GetDataGouvFrEntityByRna(rna)
	if err != nil {
		return Npo{}, err
	}
	npo := Npo(entity)

	return npo, nil
}

//goland:noinspection GoUnusedExportedFunction
func GetNpoBySiret(siret string) (Npo, error) {
	entity, err := GetDataGouvFrEntityBySiret(siret)
	if err != nil {
		return Npo{}, err
	}
	npo := Npo(entity)

	return npo, nil
}

//goland:noinspection GoUnusedExportedFunction
func SearchNposByFullText(text string) (NpoPage, error) {
	entityPage, err := GetDataGouvFrEntityPageByFullText(text)
	if err != nil {
		return NpoPage{}, err
	}
	npoPage := NpoPage{
		Count:  entityPage.Count,
		Cursor: entityPage.Cursor,
		Items:  make([]Npo, len(entityPage.Items)),
	}

	for i, s := range entityPage.Items {
		npoPage.Items[i] = Npo(s)
	}

	return npoPage, nil
}
