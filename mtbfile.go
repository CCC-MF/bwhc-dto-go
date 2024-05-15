package mtbfile

import "encoding/json"

func UnmarshalMtbFile(data []byte) (MtbFile, error) {
	var r MtbFile
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MtbFile) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MtbFile struct {
	CarePlans                     []CarePlan                     `json:"carePlans,omitempty"`
	ClaimResponses                []ClaimResponse                `json:"claimResponses,omitempty"`
	Claims                        []Claim                        `json:"claims,omitempty"`
	Consent                       Consent                        `json:"consent"`
	Diagnoses                     []Diagnosis                    `json:"diagnoses,omitempty"`
	EcogStatus                    []EcogStatus                   `json:"ecogStatus,omitempty"`
	Episode                       Episode                        `json:"episode"`
	FamilyMemberDiagnoses         []FamilyMemberDiagnosis        `json:"familyMemberDiagnoses,omitempty"`
	GeneticCounsellingRequests    []GeneticCounsellingRequest    `json:"geneticCounsellingRequests,omitempty"`
	HistologyReevaluationRequests []HistologyReevaluationRequest `json:"histologyReevaluationRequests,omitempty"`
	HistologyReports              []HistologyReport              `json:"histologyReports,omitempty"`
	LastGuidelineTherapies        []LastGuidelineTherapy         `json:"lastGuidelineTherapies,omitempty"`
	MolecularPathologyFindings    []MolecularPathologyFinding    `json:"molecularPathologyFindings,omitempty"`
	MolecularTherapies            []MolecularTherapy             `json:"molecularTherapies,omitempty"`
	NgsReports                    []NgsReport                    `json:"ngsReports,omitempty"`
	Patient                       Patient                        `json:"patient"`
	PreviousGuidelineTherapies    []PreviousGuidelineTherapy     `json:"previousGuidelineTherapies,omitempty"`
	RebiopsyRequests              []RebiopsyRequest              `json:"rebiopsyRequests,omitempty"`
	Recommendations               []Recommendation               `json:"recommendations,omitempty"`
	Responses                     []Response                     `json:"responses,omitempty"`
	Specimens                     []Specimen                     `json:"specimens,omitempty"`
	StudyInclusionRequests        []StudyInclusionRequest        `json:"studyInclusionRequests,omitempty"`
}

type CarePlan struct {
	Description               *string          `json:"description,omitempty"`
	Diagnosis                 string           `json:"diagnosis"`
	GeneticCounsellingRequest *string          `json:"geneticCounsellingRequest,omitempty"`
	ID                        string           `json:"id"`
	IssuedOn                  *string          `json:"issuedOn,omitempty"`
	NoTargetFinding           *NoTargetFinding `json:"noTargetFinding,omitempty"`
	Patient                   string           `json:"patient"`
	RebiopsyRequests          []string         `json:"rebiopsyRequests,omitempty"`
	Recommendations           []string         `json:"recommendations,omitempty"`
	StudyInclusionRequests    []string         `json:"studyInclusionRequests,omitempty"`
}

type NoTargetFinding struct {
	Diagnosis string  `json:"diagnosis"`
	IssuedOn  *string `json:"issuedOn,omitempty"`
	Patient   string  `json:"patient"`
}

type ClaimResponse struct {
	Claim    string              `json:"claim"`
	ID       string              `json:"id"`
	IssuedOn string              `json:"issuedOn"`
	Patient  string              `json:"patient"`
	Reason   *Reason             `json:"reason,omitempty"`
	Status   ClaimResponseStatus `json:"status"`
}

type Claim struct {
	ID       string `json:"id"`
	IssuedOn string `json:"issuedOn"`
	Patient  string `json:"patient"`
	Therapy  string `json:"therapy"`
}

type Consent struct {
	ID      string        `json:"id"`
	Patient string        `json:"patient"`
	Status  ConsentStatus `json:"status"`
}

type Diagnosis struct {
	GuidelineTreatmentStatus *GuidelineTreatmentStatus `json:"guidelineTreatmentStatus,omitempty"`
	HistologyResults         []string                  `json:"histologyResults,omitempty"`
	Icd10                    *DiagnosisIcd10           `json:"icd10,omitempty"`
	IcdO3T                   *IcdO3T                   `json:"icdO3T,omitempty"`
	ID                       string                    `json:"id"`
	Patient                  string                    `json:"patient"`
	RecordedOn               *string                   `json:"recordedOn,omitempty"`
	StatusHistory            []StatusHistory           `json:"statusHistory,omitempty"`
	WhoGrade                 *WhoGrade                 `json:"whoGrade,omitempty"`
}

type DiagnosisIcd10 struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type IcdO3T struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type StatusHistory struct {
	Date   string          `json:"date"`
	Status DiagnosisStatus `json:"status"`
}

type WhoGrade struct {
	Code    WHOGrade `json:"code"`
	Display *string  `json:"display,omitempty"`
	Version *string  `json:"version,omitempty"`
}

type EcogStatus struct {
	EffectiveDate *string         `json:"effectiveDate,omitempty"`
	ID            string          `json:"id"`
	Patient       string          `json:"patient"`
	Value         EcogStatusValue `json:"value"`
}

type EcogStatusValue struct {
	Code    Ecog    `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type Episode struct {
	ID      string        `json:"id"`
	Patient string        `json:"patient"`
	Period  EpisodePeriod `json:"period"`
}

type EpisodePeriod struct {
	End   *string `json:"end,omitempty"`
	Start string  `json:"start"`
}

type FamilyMemberDiagnosis struct {
	ID           string       `json:"id"`
	Patient      string       `json:"patient"`
	Relationship Relationship `json:"relationship"`
}

type Relationship struct {
	Code    CodeEnum `json:"code"`
	Display *string  `json:"display,omitempty"`
	Version *string  `json:"version,omitempty"`
}

type GeneticCounsellingRequest struct {
	ID       string  `json:"id"`
	IssuedOn *string `json:"issuedOn,omitempty"`
	Patient  string  `json:"patient"`
	Reason   string  `json:"reason"`
}

type HistologyReevaluationRequest struct {
	ID       string  `json:"id"`
	IssuedOn *string `json:"issuedOn,omitempty"`
	Patient  string  `json:"patient"`
	Specimen string  `json:"specimen"`
}

type HistologyReport struct {
	ID               string                           `json:"id"`
	IssuedOn         *string                          `json:"issuedOn,omitempty"`
	Patient          string                           `json:"patient"`
	Specimen         string                           `json:"specimen"`
	TumorCellContent *HistologyReportTumorCellContent `json:"tumorCellContent,omitempty"`
	TumorMorphology  *TumorMorphology                 `json:"tumorMorphology,omitempty"`
}

type HistologyReportTumorCellContent struct {
	ID       string                 `json:"id"`
	Method   TumorCellContentMethod `json:"method"`
	Specimen string                 `json:"specimen"`
	Value    float64                `json:"value"`
}

type TumorMorphology struct {
	ID       string               `json:"id"`
	Note     *string              `json:"note,omitempty"`
	Patient  string               `json:"patient"`
	Specimen string               `json:"specimen"`
	Value    TumorMorphologyValue `json:"value"`
}

type TumorMorphologyValue struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type LastGuidelineTherapy struct {
	Diagnosis     string                             `json:"diagnosis"`
	ID            string                             `json:"id"`
	Medication    []LastGuidelineTherapyMedication   `json:"medication,omitempty"`
	Patient       string                             `json:"patient"`
	Period        *LastGuidelineTherapyPeriod        `json:"period,omitempty"`
	ReasonStopped *LastGuidelineTherapyReasonStopped `json:"reasonStopped,omitempty"`
	TherapyLine   *int64                             `json:"therapyLine,omitempty"`
}

type LastGuidelineTherapyMedication struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  System  `json:"system"`
	Version *string `json:"version,omitempty"`
}

type LastGuidelineTherapyPeriod struct {
	End   *string `json:"end,omitempty"`
	Start string  `json:"start"`
}

type LastGuidelineTherapyReasonStopped struct {
	Code    GuidelineTherapyStopReason `json:"code"`
	Display *string                    `json:"display,omitempty"`
	Version *string                    `json:"version,omitempty"`
}

type MolecularPathologyFinding struct {
	ID                  string  `json:"id"`
	IssuedOn            *string `json:"issuedOn,omitempty"`
	Note                string  `json:"note"`
	Patient             string  `json:"patient"`
	PerformingInstitute *string `json:"performingInstitute,omitempty"`
	Specimen            string  `json:"specimen"`
}

type MolecularTherapy struct {
	History []History `json:"history"`
}

type History struct {
	BasedOn       string                 `json:"basedOn"`
	ID            string                 `json:"id"`
	NotDoneReason *NotDoneReason         `json:"notDoneReason,omitempty"`
	Note          *string                `json:"note,omitempty"`
	Patient       string                 `json:"patient"`
	RecordedOn    string                 `json:"recordedOn"`
	Status        MolecularTherapyStatus `json:"status"`
	Dosage        *Dosage                `json:"dosage,omitempty"`
	Medication    []HistoryMedication    `json:"medication,omitempty"`
	Period        *HistoryPeriod         `json:"period,omitempty"`
	ReasonStopped *HistoryReasonStopped  `json:"reasonStopped,omitempty"`
}

type HistoryMedication struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  System  `json:"system"`
	Version *string `json:"version,omitempty"`
}

type NotDoneReason struct {
	Code    NotDoneReasonCode `json:"code"`
	Display *string           `json:"display,omitempty"`
	Version *string           `json:"version,omitempty"`
}

type HistoryPeriod struct {
	End   *string `json:"end,omitempty"`
	Start string  `json:"start"`
}

type HistoryReasonStopped struct {
	Code    MolekularTherapyStopReason `json:"code"`
	Display *string                    `json:"display,omitempty"`
	Version *string                    `json:"version,omitempty"`
}

type NgsReport struct {
	Brcaness           *float64                   `json:"brcaness,omitempty"`
	CopyNumberVariants []CopyNumberVariant        `json:"copyNumberVariants,omitempty"`
	DnaFusions         []DnaFusion                `json:"dnaFusions,omitempty"`
	ID                 string                     `json:"id"`
	IssueDate          string                     `json:"issueDate"`
	Metadata           []Metadatum                `json:"metadata"`
	MSI                *float64                   `json:"msi,omitempty"`
	Patient            string                     `json:"patient"`
	RnaFusions         []RnaFusion                `json:"rnaFusions,omitempty"`
	RnaSeqs            []RnaSeq                   `json:"rnaSeqs,omitempty"`
	SequencingType     string                     `json:"sequencingType"`
	SimpleVariants     []SimpleVariant            `json:"simpleVariants,omitempty"`
	Specimen           string                     `json:"specimen"`
	Tmb                *float64                   `json:"tmb,omitempty"`
	TumorCellContent   *NgsReportTumorCellContent `json:"tumorCellContent,omitempty"`
}

type CopyNumberVariant struct {
	Chromosome            string                 `json:"chromosome"`
	CNA                   *float64               `json:"cnA,omitempty"`
	CNB                   *float64               `json:"cnB,omitempty"`
	CopyNumberNeutralLoH  []CopyNumberNeutralLoH `json:"copyNumberNeutralLoH,omitempty"`
	EndRange              EndRange               `json:"endRange"`
	ID                    string                 `json:"id"`
	RelativeCopyNumber    *float64               `json:"relativeCopyNumber,omitempty"`
	ReportedAffectedGenes []ReportedAffectedGene `json:"reportedAffectedGenes,omitempty"`
	ReportedFocality      *string                `json:"reportedFocality,omitempty"`
	StartRange            StartRange             `json:"startRange"`
	TotalCopyNumber       *int64                 `json:"totalCopyNumber,omitempty"`
	Type                  CnvType                `json:"type"`
}

type CopyNumberNeutralLoH struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type EndRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type ReportedAffectedGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type StartRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type DnaFusion struct {
	FusionPartner3Prime *DnaFusionFusionPartner3Prime `json:"fusionPartner3prime,omitempty"`
	FusionPartner5Prime *DnaFusionFusionPartner5Prime `json:"fusionPartner5prime,omitempty"`
	ID                  string                        `json:"id"`
	ReportedNumReads    *int64                        `json:"reportedNumReads,omitempty"`
}

type DnaFusionFusionPartner3Prime struct {
	Chromosome string     `json:"chromosome"`
	Gene       PurpleGene `json:"gene"`
	Position   float64    `json:"position"`
}

type PurpleGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type DnaFusionFusionPartner5Prime struct {
	Chromosome string     `json:"chromosome"`
	Gene       FluffyGene `json:"gene"`
	Position   float64    `json:"position"`
}

type FluffyGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type Metadatum struct {
	KitManufacturer string  `json:"kitManufacturer"`
	KitType         string  `json:"kitType"`
	Pipeline        *string `json:"pipeline,omitempty"`
	ReferenceGenome string  `json:"referenceGenome"`
	Sequencer       string  `json:"sequencer"`
}

type RnaFusion struct {
	CosmicID            *string                       `json:"cosmicId,omitempty"`
	Effect              *string                       `json:"effect,omitempty"`
	FusionPartner3Prime *RnaFusionFusionPartner3Prime `json:"fusionPartner3prime,omitempty"`
	FusionPartner5Prime *RnaFusionFusionPartner5Prime `json:"fusionPartner5prime,omitempty"`
	ID                  string                        `json:"id"`
	ReportedNumReads    *int64                        `json:"reportedNumReads,omitempty"`
}

type RnaFusionFusionPartner3Prime struct {
	Exon         string        `json:"exon"`
	Gene         TentacledGene `json:"gene"`
	Position     float64       `json:"position"`
	Strand       StrandEnum    `json:"strand"`
	TranscriptID string        `json:"transcriptId"`
}

type TentacledGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type RnaFusionFusionPartner5Prime struct {
	Exon         string     `json:"exon"`
	Gene         StickyGene `json:"gene"`
	Position     float64    `json:"position"`
	Strand       StrandEnum `json:"strand"`
	TranscriptID string     `json:"transcriptId"`
}

type StickyGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type RnaSeq struct {
	CohortRanking               *int64     `json:"cohortRanking,omitempty"`
	EnsemblID                   string     `json:"ensemblId"`
	EntrezID                    string     `json:"entrezId"`
	FragmentsPerKilobaseMillion float64    `json:"fragmentsPerKilobaseMillion"`
	FromNGS                     bool       `json:"fromNGS"`
	Gene                        RnaSeqGene `json:"gene"`
	ID                          string     `json:"id"`
	LibrarySize                 int64      `json:"librarySize"`
	RawCounts                   int64      `json:"rawCounts"`
	TissueCorrectedExpression   bool       `json:"tissueCorrectedExpression"`
	TranscriptID                string     `json:"transcriptId"`
}

type RnaSeqGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type SimpleVariant struct {
	AllelicFrequency float64            `json:"allelicFrequency"`
	AltAllele        string             `json:"altAllele"`
	AminoAcidChange  *AminoAcidChange   `json:"aminoAcidChange,omitempty"`
	Chromosome       string             `json:"chromosome"`
	CosmicID         *string            `json:"cosmicId,omitempty"`
	DBSNPID          *string            `json:"dbSNPId,omitempty"`
	DnaChange        *DnaChange         `json:"dnaChange,omitempty"`
	Gene             *SimpleVariantGene `json:"gene,omitempty"`
	ID               string             `json:"id"`
	Interpretation   Interpretation     `json:"interpretation"`
	ReadDepth        int64              `json:"readDepth"`
	RefAllele        string             `json:"refAllele"`
	StartEnd         StartEnd           `json:"startEnd"`
}

type AminoAcidChange struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type DnaChange struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type SimpleVariantGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}

type Interpretation struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type StartEnd struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type NgsReportTumorCellContent struct {
	ID       string                 `json:"id"`
	Method   TumorCellContentMethod `json:"method"`
	Specimen string                 `json:"specimen"`
	Value    float64                `json:"value"`
}

type Patient struct {
	BirthDate   *string `json:"birthDate,omitempty"`
	DateOfDeath *string `json:"dateOfDeath,omitempty"`
	Gender      Gender  `json:"gender"`
	ID          string  `json:"id"`
	Insurance   *string `json:"insurance,omitempty"`
	ManagingZPM *string `json:"managingZPM,omitempty"`
}

type PreviousGuidelineTherapy struct {
	Diagnosis   string                               `json:"diagnosis"`
	ID          string                               `json:"id"`
	Medication  []PreviousGuidelineTherapyMedication `json:"medication,omitempty"`
	Patient     string                               `json:"patient"`
	TherapyLine *int64                               `json:"therapyLine,omitempty"`
}

type PreviousGuidelineTherapyMedication struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  System  `json:"system"`
	Version *string `json:"version,omitempty"`
}

type RebiopsyRequest struct {
	ID       string  `json:"id"`
	IssuedOn *string `json:"issuedOn,omitempty"`
	Patient  string  `json:"patient"`
	Specimen string  `json:"specimen"`
}

type Recommendation struct {
	Diagnosis          string                     `json:"diagnosis"`
	ID                 string                     `json:"id"`
	IssuedOn           *string                    `json:"issuedOn,omitempty"`
	LevelOfEvidence    *LevelOfEvidence           `json:"levelOfEvidence,omitempty"`
	Medication         []RecommendationMedication `json:"medication,omitempty"`
	NgsReport          *string                    `json:"ngsReport,omitempty"`
	Patient            string                     `json:"patient"`
	Priority           *Priority                  `json:"priority,omitempty"`
	SupportingVariants []string                   `json:"supportingVariants,omitempty"`
}

type LevelOfEvidence struct {
	Addendums []Addendum `json:"addendums,omitempty"`
	Grading   Grading    `json:"grading"`
}

type Addendum struct {
	Code    AddendumCode `json:"code"`
	Display *string      `json:"display,omitempty"`
	Version *string      `json:"version,omitempty"`
}

type Grading struct {
	Code    GradingCode `json:"code"`
	Display *string     `json:"display,omitempty"`
	Version *string     `json:"version,omitempty"`
}

type RecommendationMedication struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  System  `json:"system"`
	Version *string `json:"version,omitempty"`
}

type Response struct {
	EffectiveDate string        `json:"effectiveDate"`
	ID            string        `json:"id"`
	Patient       string        `json:"patient"`
	Therapy       string        `json:"therapy"`
	Value         ResponseValue `json:"value"`
}

type ResponseValue struct {
	Code    Recist  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type Specimen struct {
	Collection *Collection   `json:"collection,omitempty"`
	Icd10      SpecimenIcd10 `json:"icd10"`
	ID         string        `json:"id"`
	Patient    string        `json:"patient"`
	Type       *SpecimenType `json:"type,omitempty"`
}

type Collection struct {
	Date         string           `json:"date"`
	Localization Localization     `json:"localization"`
	Method       CollectionMethod `json:"method"`
}

type SpecimenIcd10 struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	Version *string `json:"version,omitempty"`
}

type StudyInclusionRequest struct {
	ID        string  `json:"id"`
	IssuedOn  *string `json:"issuedOn,omitempty"`
	NctNumber string  `json:"nctNumber"`
	Patient   string  `json:"patient"`
	Reason    string  `json:"reason"`
}

type Reason string

const (
	InsufficientEvidence        Reason = "insufficient-evidence"
	ReasonOther                 Reason = "other"
	StandardTherapyNotExhausted Reason = "standard-therapy-not-exhausted"
)

type ClaimResponseStatus string

const (
	Accepted                    ClaimResponseStatus = "accepted"
	ClaimResponseStatusRejected ClaimResponseStatus = "rejected"
)

type ConsentStatus string

const (
	Active                ConsentStatus = "active"
	ConsentStatusRejected ConsentStatus = "rejected"
)

type GuidelineTreatmentStatus string

const (
	Exhausted                       GuidelineTreatmentStatus = "exhausted"
	GuidelineTreatmentStatusUnknown GuidelineTreatmentStatus = "unknown"
	Impossible                      GuidelineTreatmentStatus = "impossible"
	NoGuidelinesAvailable           GuidelineTreatmentStatus = "no-guidelines-available"
	NonExhausted                    GuidelineTreatmentStatus = "non-exhausted"
)

type DiagnosisStatus string

const (
	DiagnosisStatusUnknown DiagnosisStatus = "unknown"
	Local                  DiagnosisStatus = "local"
	Metastasized           DiagnosisStatus = "metastasized"
	TumorFree              DiagnosisStatus = "tumor-free"
)

type WHOGrade string

const (
	I   WHOGrade = "I"
	Ii  WHOGrade = "II"
	Iii WHOGrade = "III"
	Iv  WHOGrade = "IV"
)

type Ecog string

const (
	Ecog1 Ecog = "1"
	Ecog2 Ecog = "2"
	Ecog3 Ecog = "3"
	Ecog4 Ecog = "4"
	The0  Ecog = "0"
)

type CodeEnum string

const (
	EXT     CodeEnum = "EXT"
	Fammemb CodeEnum = "FAMMEMB"
)

type TumorCellContentMethod string

const (
	Bioinformatic TumorCellContentMethod = "bioinformatic"
	Histologic    TumorCellContentMethod = "histologic"
)

type System string

const (
	Atc          System = "ATC"
	Unregistered System = "Unregistered"
)

type GuidelineTherapyStopReason string

const (
	ChronicRemission                        GuidelineTherapyStopReason = "chronic-remission"
	GuidelineTherapyStopReasonDeterioration GuidelineTherapyStopReason = "deterioration"
	GuidelineTherapyStopReasonOther         GuidelineTherapyStopReason = "other"
	GuidelineTherapyStopReasonPatientWish   GuidelineTherapyStopReason = "patient-wish"
	GuidelineTherapyStopReasonProgression   GuidelineTherapyStopReason = "progression"
	GuidelineTherapyStopReasonToxicity      GuidelineTherapyStopReason = "toxicity"
	GuidelineTherapyStopReasonUnknown       GuidelineTherapyStopReason = "unknown"
)

type Dosage string

const (
	GreaterOrEqual50 Dosage = ">=50%"
	Less50           Dosage = "<50%"
)

type NotDoneReasonCode string

const (
	LostToFu                         NotDoneReasonCode = "lost-to-fu"
	NoIndication                     NotDoneReasonCode = "no-indication"
	NotDoneReasonContinuedExternally NotDoneReasonCode = "continued-externally"
	NotDoneReasonMedicalReason       NotDoneReasonCode = "medical-reason"
	NotDoneReasonOther               NotDoneReasonCode = "other"
	NotDoneReasonOtherTherapyChosen  NotDoneReasonCode = "other-therapy-chosen"
	NotDoneReasonPatientDeath        NotDoneReasonCode = "patient-death"
	NotDoneReasonUnknown             NotDoneReasonCode = "unknown"
	PatientRefusal                   NotDoneReasonCode = "patient-refusal"
	PaymentPending                   NotDoneReasonCode = "payment-pending"
	PaymentRefused                   NotDoneReasonCode = "payment-refused"
)

type MolekularTherapyStopReason string

const (
	MolekularTherapyStopReasonContinuedExternally MolekularTherapyStopReason = "continued-externally"
	MolekularTherapyStopReasonDeterioration       MolekularTherapyStopReason = "deterioration"
	MolekularTherapyStopReasonMedicalReason       MolekularTherapyStopReason = "medical-reason"
	MolekularTherapyStopReasonOther               MolekularTherapyStopReason = "other"
	MolekularTherapyStopReasonOtherTherapyChosen  MolekularTherapyStopReason = "other-therapy-chosen"
	MolekularTherapyStopReasonPatientDeath        MolekularTherapyStopReason = "patient-death"
	MolekularTherapyStopReasonPatientWish         MolekularTherapyStopReason = "patient-wish"
	MolekularTherapyStopReasonProgression         MolekularTherapyStopReason = "progression"
	MolekularTherapyStopReasonToxicity            MolekularTherapyStopReason = "toxicity"
	MolekularTherapyStopReasonUnknown             MolekularTherapyStopReason = "unknown"
	PaymentEnded                                  MolekularTherapyStopReason = "payment-ended"
	Remission                                     MolekularTherapyStopReason = "remission"
)

type MolecularTherapyStatus string

const (
	Completed MolecularTherapyStatus = "completed"
	NotDone   MolecularTherapyStatus = "not-done"
	OnGoing   MolecularTherapyStatus = "on-going"
	Stopped   MolecularTherapyStatus = "stopped"
)

type CnvType string

const (
	HighLevelGain CnvType = "high-level-gain"
	Loss          CnvType = "loss"
	LowLevelGain  CnvType = "low-level-gain"
)

type StrandEnum string

const (
	Empty  StrandEnum = "+"
	Strand StrandEnum = "-"
)

type Gender string

const (
	Female        Gender = "female"
	GenderOther   Gender = "other"
	GenderUnknown Gender = "unknown"
	Male          Gender = "male"
)

type AddendumCode string

const (
	AddendumIv AddendumCode = "iv"
	Is         AddendumCode = "is"
	R          AddendumCode = "R"
	Z          AddendumCode = "Z"
)

type GradingCode string

const (
	M1A GradingCode = "m1A"
	M1B GradingCode = "m1B"
	M1C GradingCode = "m1C"
	M2A GradingCode = "m2A"
	M2B GradingCode = "m2B"
	M2C GradingCode = "m2C"
	M3  GradingCode = "m3"
	M4  GradingCode = "m4"
)

type Priority string

const (
	Priority1 Priority = "1"
	Priority2 Priority = "2"
	Priority3 Priority = "3"
	Priority4 Priority = "4"
)

type Recist string

const (
	CR  Recist = "CR"
	Mr  Recist = "MR"
	Na  Recist = "NA"
	Nya Recist = "NYA"
	PD  Recist = "PD"
	PR  Recist = "PR"
	SD  Recist = "SD"
)

type Localization string

const (
	LocalizationUnknown Localization = "unknown"
	Metastasis          Localization = "metastasis"
	PrimaryTumor        Localization = "primary-tumor"
)

type CollectionMethod string

const (
	Biopsy                       CollectionMethod = "biopsy"
	CollectionMethodLiquidBiopsy CollectionMethod = "liquid-biopsy"
	CollectionMethodUnknown      CollectionMethod = "unknown"
	Cytology                     CollectionMethod = "cytology"
	Resection                    CollectionMethod = "resection"
)

type SpecimenType string

const (
	CryoFrozen               SpecimenType = "cryo-frozen"
	Ffpe                     SpecimenType = "FFPE"
	FreshTissue              SpecimenType = "fresh-tissue"
	SpecimenTypeLiquidBiopsy SpecimenType = "liquid-biopsy"
	SpecimenTypeUnknown      SpecimenType = "unknown"
)
