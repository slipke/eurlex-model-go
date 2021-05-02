package model

import "fmt"

//
// Notice
//

type Notice struct {
	Decoding      string           `xml:"decoding,attr"`
	Type          string           `xml:"type,attr"`
	Work          *Work            `xml:"WORK"`
	Inverse       []*Inverse       `xml:"INVERSE"`
	Expression    []*Expression    `xml:"EXPRESSION"`
	Manifestation []*Manifestation `xml:"MANIFESTATION"`
}

func (n *Notice) String() string {
	return fmt.Sprintf(
		"Notice (decoding: %s, type: %s), work = %p, inverse = %p, expression = %p, manifestation = %p",
		n.Decoding,
		n.Type,
		n.Work,
		n.Inverse,
		n.Expression,
		n.Manifestation,
	)
}

//
// Work
//

type Work struct {
	Identifier
	ResourceLegalIsAboutSubjectMatter              []*ResourceLegalIsAboutSubjectMatter `xml:"RESOURCE_LEGAL_IS_ABOUT_SUBJECT-MATTER"`
	WorkCitesWork                                  []*Link                              `xml:"WORK_CITES_WORK"`
	WorkHasExpression                              []*Link                              `xml:"WORK_HAS_EXPRESSION"` // Link to Expressions
	ResourceLegalELI                               string                               `xml:"RESOURCE_LEGAL_ELI>VALUE"`
	ResourceLegalImplicitlyRepealsResourceLegal    []*Link                              `xml:"RESOURCE_LEGAL_IMPLICITLY_REPEALS_RESOURCE_LEGAL"`
	Date                                           []*Date                              `xml:"DATE"`
	IdentifierItem                                 []string                             `xml:"IDENTIFIER>VALUE"`
	IsAbout                                        []*Concept                           `xml:"IS_ABOUT"`
	ResourceLegalNumberNatural                     string                               `xml:"RESOURCE_LEGAL_NUMBER_NATURAL>VALUE"`
	CreationDate                                   *Date                                `xml:"CREATIONDATE"`
	ResourceLegalInformationMiscellaneous          string                               `xml:"RESOURCE_LEGAL_INFORMATION_MISCELLANEOUS>VALUE"`
	WorkIsAboutConceptEurovoc                      []*ConceptFacet                      `xml:"WORK_IS_ABOUT_CONCEPT_EUROVOC"`
	ResourceLegalHasTypeActConceptTypeAct          *Concept                             `xml:"RESOURCE_LEGAL_HAS_TYPE_ACT_CONCEPT_TYPE_ACT"`
	ResourceLegalYear                              string                               `xml:"RESOURCE_LEGAL_YEAR>VALUE"`
	BasedOn                                        []*Link                              `xml:"BASED_ON"`
	ResourceLegalDateEntryIntoForce                []*Date                              `xml:"RESOURCE_LEGAL_DATE_ENTRY-INTO-FORCE"`
	ResourceLegalRepealsResourceLegal              []*Link                              `xml:"RESOURCE_LEGAL_REPEALS_RESOURCE_LEGAL"`
	ResourceLegalAmendsResourceLegal               []*Link                              `xml:"RESOURCE_LEGAL_AMENDS_RESOURCE_LEGAL"`
	IDCelex                                        string                               `xml:"ID_CELEX>VALUE"`
	WorkDateDocument                               *Date                                `xml:"WORK_DATE_DOCUMENT"`
	WorkCreatedByAgent                             []*Concept                           `xml:"WORK_CREATED_BY_AGENT"`
	ResourceLegalPublishedInOfficialJournal        []*Link                              `xml:"RESOURCE_LEGAL_PUBLISHED_IN_OFFICIAL-JOURNAL"`
	DateDocument                                   *Date                                `xml:"DATE_DOCUMENT"`
	ResourceLegalBasedOnConceptTreaty              *Concept                             `xml:"RESOURCE_LEGAL_BASED_ON_CONCEPT_TREATY"`
	ELI                                            string                               `xml:"ELI>VALUE"`
	ResouceLegalType                               string                               `xml:"RESOURCE_LEGAL_TYPE>VALUE"`
	ResouceLegalDateEndOfValidity                  *Date                                `xml:"RESOURCE_LEGAL_DATE_END-OF-VALIDITY"`
	WorkIDDocument                                 []string                             `xml:"WORK_ID_DOCUMENT>VALUE"`
	ResourceLegalInForce                           bool                                 `xml:"RESOURCE_LEGAL_IN-FORCE>VALUE"`
	ResourceLegalBasedOnResourceLegal              []*Link                              `xml:"RESOURCE_LEGAL_BASED_ON_RESOURCE_LEGAL"`
	ResourceLegalRepertoire                        string                               `xml:"RESOURCE_LEGAL_REPERTOIRE>VALUE"`
	IDSector                                       string                               `xml:"ID_SECTOR>VALUE"`
	DateCreationLegacy                             *Date                                `xml:"DATE_CREATION_LEGACY"`
	ResourceLegalDerogatedByResourceLegal          *Link                                `xml:"RESOURCE_LEGAL_DEROGATED_BY_RESOURCE_LEGAL"`
	ResourceLegalNumberNaturalCelex                string                               `xml:"RESOURCE_LEGAL_NUMBER_NATURAL_CELEX>VALUE"`
	WorkHasResourceType                            *Concept                             `xml:"WORK_HAS_RESOURCE-TYPE"`
	ResourceLegalIDCelex                           string                               `xml:"RESOURCE_LEGAL_ID_CELEX>VALUE"`
	ResourceLegalCommentInternal                   string                               `xml:"RESOURCE_LEGAL_COMMENT_INTERNAL>VALUE"`
	CreatedBy                                      []*Concept                           `xml:"CREATED_BY"`
	ResourceLegalPublishedInSpecialOfficialJournal *Link                                `xml:"RESOURCE_LEGAL_PUBLISHED_IN_SPECIAL-OFFICIAL-JOURNAL"`
	WorkPartOfCollectionDocument                   []*Concept                           `xml:"WORK_PART_OF_COLLECTION_DOCUMENT"`
	WorkLegalAdoptsResourceLegal                   []*Link                              `xml:"RESOURCE_LEGAL_ADOPTS_RESOURCE_LEGAL"`
	WorkTableOfContents                            string                               `xml:"WORK_TABLE-OF-CONTENTS>VALUE"`
	LastModificationDate                           *Date                                `xml:"LASTMODIFICATIONDATE"`
	Type                                           []string                             `xml:"TYPE"`

	// From eurlex-ws-go
	DatePublication                         string `xml:"DATE_PUBLICATION>VALUE"`
	OfficialJournalClass                    string `xml:"OFFICIAL-JOURNAL_CLASS>VALUE"`
	OfficialJournalNumber                   string `xml:"OFFICIAL-JOURNAL_NUMBER>VALUE"`
	OfficialJournalPartOfCollectionDocument *Link  `xml:"OFFICIAL-JOURNAL_PART_OF_COLLECTION_DOCUMENT"`
	OfficialJournalYear                     string `xml:"OFFICIAL-JOURNAL_YEAR>VALUE"`

	// @TODO Same as is INVERSE
	ResourceLegalBasisForActConsolidated          []*Link `xml:"RESOURCE_LEGAL_BASIS_FOR_ACT_CONSOLIDATED"`
	ResourceLegalConsolidatedByActConsolidated    []*Link `xml:"RESOURCE_LEGAL_CONSOLIDATED_BY_ACT_CONSOLIDATED"`
	ResourceLegalCorrectedByResourceLegal         []*Link `xml:"RESOURCE_LEGAL_CORRECTED_BY_RESOURCE_LEGAL"`
	ResourceLegalAmendedByResourceLegal           []*Link `xml:"RESOURCE_LEGAL_AMENDED_BY_RESOURCE_LEGAL"`
	ResourceLegalInterpretationRequestedByCaseLaw []*Link `xml:"RESOURCE_LEGAL_INTERPRETATION_REQUESTED_BY_CASE-LAW"`
	WorkSummarizedBySummary                       *Link   `xml:"WORK_SUMMARIZED_BY_SUMMARY"`
	ResourceLegalProducedByDossier                []*Link `xml:"RESOURCE_LEGAL_PRODUCED_BY_DOSSIER"`
	ResourceLegalAmendmentProposedByResourceLegal []*Link `xml:"RESOURCE_LEGAL_AMENDMENT_PROPOSED_BY_RESOURCE_LEGAL"`
	WorkPartOfDossier                             []*Link `xml:"WORK_PART_OF_DOSSIER"`
	WorkPartOfEventLegal                          []*Link `xml:"WORK_PART_OF_EVENT_LEGAL"`

	// ResourceLegalIsAboutConceptDirectoryCode *? `xml:"RESOURCE_LEGAL_IS_ABOUT_CONCEPT_DIRECTORY-CODE"` // @TODO Fix weird structure
}

func (w *Work) String() string {
	u := ""
	if w.URI != nil {
		u = w.URI.Value
	}

	return fmt.Sprintf(
		"Work URI = %s, len(ResourceLegalIsAboutSubjectMatter) = %d, len(WorkCitesWork) = %d, len(WorkHasExpression) = %d, Identifier = %+v, CreationDate = %+v, IDCelex  %s, ELI = %s, IDSector = %s, Type = %+v",
		u,
		len(w.ResourceLegalIsAboutSubjectMatter),
		len(w.WorkCitesWork),
		len(w.WorkHasExpression),
		w.Identifier,
		w.CreationDate,
		w.IDCelex,
		w.ELI,
		w.IDSector,
		w.Type,
	)
}

// @TODO Rename to concept_level ??
// @TODO Fix weird structure
// type="concept_level"
type ResourceLegalIsAboutSubjectMatter struct {
	Type                               string   `xml:"type,attr"`
	ResourceLegalIsAboutSubjectMatter1 *Concept `xml:"RESOURCE_LEGAL_IS_ABOUT_SUBJECT-MATTER_1"`
}

//
// Inverse
//

type Inverse struct {
	Complete                                      bool    `xml:"complete,attr"`
	ManifestationType                             string  `xml:"manifestation-type,attr"`
	ResourceLegalBasisForActConsolidated          []*Link `xml:"RESOURCE_LEGAL_BASIS_FOR_ACT_CONSOLIDATED"`
	ResourceLegalConsolidatedByActConsolidated    []*Link `xml:"RESOURCE_LEGAL_CONSOLIDATED_BY_ACT_CONSOLIDATED"`
	ResourceLegalCorrectedByResourceLegal         []*Link `xml:"RESOURCE_LEGAL_CORRECTED_BY_RESOURCE_LEGAL"`
	ResourceLegalAmendedByResourceLegal           []*Link `xml:"RESOURCE_LEGAL_AMENDED_BY_RESOURCE_LEGAL"`
	ResourceLegalInterpretationRequestedByCaseLaw []*Link `xml:"RESOURCE_LEGAL_INTERPRETATION_REQUESTED_BY_CASE-LAW"`
	WorkSummarizedBySummary                       *Link   `xml:"WORK_SUMMARIZED_BY_SUMMARY"`
	ResourceLegalProducedByDossier                []*Link `xml:"RESOURCE_LEGAL_PRODUCED_BY_DOSSIER"`
	ResourceLegalAmendmentProposedByResourceLegal []*Link `xml:"RESOURCE_LEGAL_AMENDMENT_PROPOSED_BY_RESOURCE_LEGAL"`
	WorkPartOfDossier                             []*Link `xml:"WORK_PART_OF_DOSSIER"`
	WorkPartOfEventLegal                          []*Link `xml:"WORK_PART_OF_EVENT_LEGAL"`
}

func (i *Inverse) String() string {
	return fmt.Sprintf(
		"Inverse (Complete = %t, ManifestationType = %v) %d %d %d %d %d %d %d %d %d",
		i.Complete,
		i.ManifestationType,
		len(i.ResourceLegalBasisForActConsolidated),
		len(i.ResourceLegalConsolidatedByActConsolidated),
		len(i.ResourceLegalCorrectedByResourceLegal),
		len(i.ResourceLegalAmendedByResourceLegal),
		len(i.ResourceLegalInterpretationRequestedByCaseLaw),
		len(i.ResourceLegalProducedByDossier),
		len(i.ResourceLegalAmendmentProposedByResourceLegal),
		len(i.WorkPartOfDossier),
		len(i.WorkPartOfEventLegal),
	)
}

//
// Expression
//

type Expression struct {
	Identifier
	ExpressionManifestedByManifestation []*Link  `xml:"EXPRESSION_MANIFESTED_BY_MANIFESTATION"` // List of manifestations for expression (^= different languages for current expression)
	Lang                                []string `xml:"LANG>VALUE"`
	LastModificationDate                *Date    `xml:"LASTMODIFICATIONDATE"`
	Title                               []string `xml:"TITLE>VALUE"`
	ExpressionTitle                     string   `xml:"EXPRESSION_TITLE>VALUE"`
	CreationDate                        *Date    `xml:"CREATIONDATE"`
	ExpressionUsesLanguage              *Concept `xml:"EXPRESSION_USES_LANGUAGE"`
	ExpressionBelongsToWork             *Link    `xml:"EXPRESSION_BELONGS_TO_WORK"` // Link to work
	ExpressionTitleShort                string   `xml:"EXPRESSION_TITLE_SHORT>VALUE"`
}

func (e *Expression) String() string {
	return fmt.Sprintf(
		"Expression URI = %+v, Lang = %+v, CreationDate = %+v, LastModificationDate = %+v, ExpressionTitle = %s, ExpressionTitleShort = %s, len(ExpressionManifestedByManifestation) = %d",
		e.URI,
		e.Lang,
		e.CreationDate,
		e.LastModificationDate,
		e.ExpressionTitle,
		e.ExpressionTitleShort,
		len(e.ExpressionManifestedByManifestation),
	)
}

//
// Manifestation
//

type Manifestation struct {
	ManifestationTypeAttr string `xml:"manifestation-type,attr"`
	Identifier
	LastModificationDate                           *Date                   `xml:"LASTMODIFICATIONDATE"`
	ManifestationManifestsExpression               *Link                   `xml:"MANIFESTATION_MANIFESTS_EXPRESSION"`
	CreationDate                                   *Date                   `xml:"CREATIONDATE"`
	ManifestationHasItem                           []*ManifestationHasItem `xml:"MANIFESTATION_HAS_ITEM"` // Attachments
	ManifestationType                              string                  `xml:"MANIFESTATION_TYPE>VALUE"`
	ManifestationSourcePool                        string                  `xml:"MANIFESTATION_SOURCE-POOL>VALUE"`
	Type                                           string                  `xml:"TYPE"`
	ManifestationPartOfManifestation               *Link                   `xml:"MANIFESTATION_PART_OF_MANIFESTATION"`
	AuthorPrinter                                  []string                `xml:"AUTHOR_PRINTER>VALUE"`
	ManifestationOfficialJournalPartSectionOJ      string                  `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_SECTION_OJ>VALUE"`
	ManifestationOfficialJournalPartPageFirst      string                  `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_PAGE_FIRST>VALUE"`
	PageLast                                       string                  `xml:"PAGE_LAST>VALUE"`
	ManifestationOfficialJournalPartAuthorPrinter  []string                `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_AUTHOR_PRINTER>VALUE"`
	TypedocPrinter                                 string                  `xml:"TYPEDOC_PRINTER>VALUE"`
	PageFirst                                      string                  `xml:"PAGE_FIRST>VALUE"`
	ManifestationOfficialJournalPartTypedocPrinter string                  `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_TYPEDOC_PRINTER>VALUE"`
	ManifestationOfficialJournalPartPageLast       string                  `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_PAGE_LAST>VALUE"`
	ManifestationOfficialJournalPartPagesTotal     string                  `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_PAGES_TOTAL>VALUE"`
}

func (m *Manifestation) String() string {
	return fmt.Sprintf(
		"Manifestation (manifestation-type = %s) URI = %s, LastModificationDate = %+v, CreationDate = %+v, len(ManifestationHasItem) = %d, type = %s",
		m.ManifestationTypeAttr,
		m.URI,
		m.LastModificationDate,
		m.CreationDate,
		len(m.ManifestationHasItem),
		m.Type,
	)
}

type ManifestationHasItem struct {
	Identifier
	TechMD                     *TechMD `xml:"TECHMD"`
	ItemIdentifier             string  `xml:"ITEM_IDENTIFIER>VALUE"`
	IdentifierItem             string  `xml:"IDENTIFIER>VALUE"`
	BelongsTo                  *Link   `xml:"BELONGS_TO"`
	ItemBelongsToManifestation *Link   `xml:"ITEM_BELONGS_TO_MANIFESTATION"`
}

type TechMD struct {
	ManifestationType      string `xml:"MANIFESTATION_TYPE"`
	MimeType               string `xml:"MIME-TYPE"`
	StreamCompositionLevel string `xml:"STREAM_COMPOSITION_LEVEL>VALUE"`
	StreamSize             string `xml:"STREAM_SIZE>VALUE"`
	StreamName             string `xml:"STREAM_NAME>VALUE"`
	StreamLabel            string `xml:"STREAM_LABEL>VALUE"`
	StreamOrder            string `xml:"STREAM_ORDER>VALUE"`
}

//
// Utility / General
//

// Date encapsulates VALUE, YEAR, MONTH and DAY
// type="date"
type Date struct {
	Type       string        `xml:"type,attr"`
	Value      string        `xml:"VALUE"`
	Year       string        `xml:"YEAR"`
	Month      string        `xml:"MONTH"`
	Day        string        `xml:"DAY"`
	Annotation []*Annotation `xml:"ANNOTATION"`
}

type Identifier struct {
	URI    *URI      `xml:"URI"`
	SameAs []*SameAs `xml:"SAMEAS"`
}

// Link encapsulates properties URI and SAMEAS and attribute type
// type="link"
type Link struct {
	Type string `xml:"type,attr"`
	Identifier
	EmbeddedNotice *Notice       `xml:"EMBEDDED_NOTICE"`
	Annotation     []*Annotation `xml:"ANNOTATION"`
	IdentifierVal  string        `xml:"IDENTIFIER"`
	PrefLabel      string        `xml:"PREFLABEL"`
	AltLabel       []string      `xml:"ALTLABEL"`
}

type URI struct {
	Value      string `xml:"VALUE"`
	Identifier string `xml:"IDENTIFIER"`
	Type       string `xml:"TYPE"`
}

type SameAs struct {
	URI *URI `xml:"URI"`
}

// @TODO Split annotation?
type Annotation struct {
	// RESOURCE_LEGAL_CORRECTED_BY_RESOURCE_LEGAL
	BuildInfo        string `xml:"BUILD_INFO"`
	TypeOfLinkTarget string `xml:"TYPE_OF_LINK_TARGET"`
	LanguageList     string `xml:"LANGUAGE_LIST"`
	// RESOURCE_LEGAL_AMENDED_BY_RESOURCE_LEGAL
	ReferenceToModifiedLocation string `xml:"REFERENCE_TO_MODIFIED_LOCATION"`
	Role2                       string `xml:"ROLE2"`
	StartOfValidity             string `xml:"START_OF_VALIDITY"`
	// RESOURCE_LEGAL_DATE_ENTRY-INTO-FORCE
	CommentOnDate string `xml:"COMMENT_ON_DATE"`
	TypeOfDate    string `xml:"TYPE_OF_DATE"`
}

// Concept .
// type="concept"
type Concept struct {
	Type       string   `xml:"type,attr"`
	URI        *URI     `xml:"URI"`
	OpCode     string   `xml:"OP-CODE"`
	Identifier string   `xml:"IDENTIFIER"`
	PrefLabel  string   `xml:"PREFLABEL"`
	AltLabel   []string `xml:"ALTLABEL"`
	CompactURI string   `xml:"COMPACT_URI"`
	Fallback   string   `xml:"FALLBACK"`
}

// ConceptFacet .
// type="concept_facet"
// @TODO Complex datatype, children: <property-name_{facet.acronym} type="concept">...
type ConceptFacet struct {
	Type                             string   `xml:"type,attr"`
	WorkIsAboutConceptEurovocConcept *Concept `xml:"WORK_IS_ABOUT_CONCEPT_EUROVOC_CONCEPT"`
	WorkIsAboutConceptEurovocTT      *Concept `xml:"WORK_IS_ABOUT_CONCEPT_EUROVOC_TT"`
	WorkIsAboutConceptEurovocDOM     *Concept `xml:"WORK_IS_ABOUT_CONCEPT_EUROVOC_DOM"`
	WorkIsAboutConceptEurovocMTH     *Concept `xml:"WORK_IS_ABOUT_CONCEPT_EUROVOC_MTH"`
}
