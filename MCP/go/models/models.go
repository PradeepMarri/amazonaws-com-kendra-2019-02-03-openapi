package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// InlineCustomDocumentEnrichmentConfiguration represents the InlineCustomDocumentEnrichmentConfiguration schema from the OpenAPI specification
type InlineCustomDocumentEnrichmentConfiguration struct {
	Condition interface{} `json:"Condition,omitempty"`
	Documentcontentdeletion interface{} `json:"DocumentContentDeletion,omitempty"`
	Target interface{} `json:"Target,omitempty"`
}

// SuggestionValue represents the SuggestionValue schema from the OpenAPI specification
type SuggestionValue struct {
	Text interface{} `json:"Text,omitempty"`
}

// DisassociatePersonasFromEntitiesResponse represents the DisassociatePersonasFromEntitiesResponse schema from the OpenAPI specification
type DisassociatePersonasFromEntitiesResponse struct {
	Failedentitylist interface{} `json:"FailedEntityList,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tagkeys interface{} `json:"TagKeys"`
}

// CreateThesaurusResponse represents the CreateThesaurusResponse schema from the OpenAPI specification
type CreateThesaurusResponse struct {
	Id interface{} `json:"Id,omitempty"`
}

// CreateExperienceRequest represents the CreateExperienceRequest schema from the OpenAPI specification
type CreateExperienceRequest struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
}

// ListExperienceEntitiesRequest represents the ListExperienceEntitiesRequest schema from the OpenAPI specification
type ListExperienceEntitiesRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateDataSourceRequest represents the UpdateDataSourceRequest schema from the OpenAPI specification
type UpdateDataSourceRequest struct {
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Customdocumentenrichmentconfiguration interface{} `json:"CustomDocumentEnrichmentConfiguration,omitempty"`
	Id interface{} `json:"Id"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name,omitempty"`
}

// GroupMembers represents the GroupMembers schema from the OpenAPI specification
type GroupMembers struct {
	Membergroups interface{} `json:"MemberGroups,omitempty"`
	Memberusers interface{} `json:"MemberUsers,omitempty"`
	S3pathforgroupmembers interface{} `json:"S3PathforGroupMembers,omitempty"`
}

// DocumentsMetadataConfiguration represents the DocumentsMetadataConfiguration schema from the OpenAPI specification
type DocumentsMetadataConfiguration struct {
	S3prefix interface{} `json:"S3Prefix,omitempty"`
}

// ListDataSourceSyncJobsRequest represents the ListDataSourceSyncJobsRequest schema from the OpenAPI specification
type ListDataSourceSyncJobsRequest struct {
	Statusfilter interface{} `json:"StatusFilter,omitempty"`
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Starttimefilter interface{} `json:"StartTimeFilter,omitempty"`
}

// GitHubConfiguration represents the GitHubConfiguration schema from the OpenAPI specification
type GitHubConfiguration struct {
	Inclusionfilenamepatterns interface{} `json:"InclusionFileNamePatterns,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Githubissuedocumentconfigurationfieldmappings interface{} `json:"GitHubIssueDocumentConfigurationFieldMappings,omitempty"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Githubpullrequestcommentconfigurationfieldmappings interface{} `json:"GitHubPullRequestCommentConfigurationFieldMappings,omitempty"`
	Githubrepositoryconfigurationfieldmappings interface{} `json:"GitHubRepositoryConfigurationFieldMappings,omitempty"`
	Inclusionfoldernamepatterns interface{} `json:"InclusionFolderNamePatterns,omitempty"`
	Githubdocumentcrawlproperties interface{} `json:"GitHubDocumentCrawlProperties,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Githubpullrequestdocumentattachmentconfigurationfieldmappings interface{} `json:"GitHubPullRequestDocumentAttachmentConfigurationFieldMappings,omitempty"`
	Githubissuecommentconfigurationfieldmappings interface{} `json:"GitHubIssueCommentConfigurationFieldMappings,omitempty"`
	Repositoryfilter interface{} `json:"RepositoryFilter,omitempty"`
	Inclusionfiletypepatterns interface{} `json:"InclusionFileTypePatterns,omitempty"`
	Saasconfiguration interface{} `json:"SaaSConfiguration,omitempty"`
	Exclusionfilenamepatterns interface{} `json:"ExclusionFileNamePatterns,omitempty"`
	Githubissueattachmentconfigurationfieldmappings interface{} `json:"GitHubIssueAttachmentConfigurationFieldMappings,omitempty"`
	Githubcommitconfigurationfieldmappings interface{} `json:"GitHubCommitConfigurationFieldMappings,omitempty"`
	Githubpullrequestdocumentconfigurationfieldmappings interface{} `json:"GitHubPullRequestDocumentConfigurationFieldMappings,omitempty"`
	Exclusionfiletypepatterns interface{} `json:"ExclusionFileTypePatterns,omitempty"`
	Onpremiseconfiguration interface{} `json:"OnPremiseConfiguration,omitempty"`
	Exclusionfoldernamepatterns interface{} `json:"ExclusionFolderNamePatterns,omitempty"`
	Usechangelog interface{} `json:"UseChangeLog,omitempty"`
}

// Suggestion represents the Suggestion schema from the OpenAPI specification
type Suggestion struct {
	Id interface{} `json:"Id,omitempty"`
	Sourcedocuments interface{} `json:"SourceDocuments,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// ServerSideEncryptionConfiguration represents the ServerSideEncryptionConfiguration schema from the OpenAPI specification
type ServerSideEncryptionConfiguration struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// FsxConfiguration represents the FsxConfiguration schema from the OpenAPI specification
type FsxConfiguration struct {
	Vpcconfiguration interface{} `json:"VpcConfiguration"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Filesystemid interface{} `json:"FileSystemId"`
	Filesystemtype interface{} `json:"FileSystemType"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Secretarn interface{} `json:"SecretArn,omitempty"`
}

// SalesforceStandardKnowledgeArticleTypeConfiguration represents the SalesforceStandardKnowledgeArticleTypeConfiguration schema from the OpenAPI specification
type SalesforceStandardKnowledgeArticleTypeConfiguration struct {
	Documentdatafieldname interface{} `json:"DocumentDataFieldName"`
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
}

// ConnectionConfiguration represents the ConnectionConfiguration schema from the OpenAPI specification
type ConnectionConfiguration struct {
	Databasehost interface{} `json:"DatabaseHost"`
	Databasename interface{} `json:"DatabaseName"`
	Databaseport interface{} `json:"DatabasePort"`
	Secretarn interface{} `json:"SecretArn"`
	Tablename interface{} `json:"TableName"`
}

// UpdateExperienceRequest represents the UpdateExperienceRequest schema from the OpenAPI specification
type UpdateExperienceRequest struct {
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name,omitempty"`
}

// DescribeExperienceRequest represents the DescribeExperienceRequest schema from the OpenAPI specification
type DescribeExperienceRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// WebCrawlerConfiguration represents the WebCrawlerConfiguration schema from the OpenAPI specification
type WebCrawlerConfiguration struct {
	Proxyconfiguration interface{} `json:"ProxyConfiguration,omitempty"`
	Urlexclusionpatterns interface{} `json:"UrlExclusionPatterns,omitempty"`
	Urlinclusionpatterns interface{} `json:"UrlInclusionPatterns,omitempty"`
	Maxlinksperpage interface{} `json:"MaxLinksPerPage,omitempty"`
	Authenticationconfiguration interface{} `json:"AuthenticationConfiguration,omitempty"`
	Crawldepth interface{} `json:"CrawlDepth,omitempty"`
	Maxcontentsizeperpageinmegabytes interface{} `json:"MaxContentSizePerPageInMegaBytes,omitempty"`
	Maxurlsperminutecrawlrate interface{} `json:"MaxUrlsPerMinuteCrawlRate,omitempty"`
	Urls interface{} `json:"Urls"`
}

// Facet represents the Facet schema from the OpenAPI specification
type Facet struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Documentattributekey interface{} `json:"DocumentAttributeKey,omitempty"`
	Facets interface{} `json:"Facets,omitempty"`
}

// DescribeIndexResponse represents the DescribeIndexResponse schema from the OpenAPI specification
type DescribeIndexResponse struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Usergroupresolutionconfiguration interface{} `json:"UserGroupResolutionConfiguration,omitempty"`
	Documentmetadataconfigurations interface{} `json:"DocumentMetadataConfigurations,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Edition interface{} `json:"Edition,omitempty"`
	Serversideencryptionconfiguration interface{} `json:"ServerSideEncryptionConfiguration,omitempty"`
	Capacityunits interface{} `json:"CapacityUnits,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Indexstatistics interface{} `json:"IndexStatistics,omitempty"`
	Usercontextpolicy interface{} `json:"UserContextPolicy,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Usertokenconfigurations interface{} `json:"UserTokenConfigurations,omitempty"`
}

// BatchGetDocumentStatusRequest represents the BatchGetDocumentStatusRequest schema from the OpenAPI specification
type BatchGetDocumentStatusRequest struct {
	Documentinfolist interface{} `json:"DocumentInfoList"`
	Indexid interface{} `json:"IndexId"`
}

// UserGroupResolutionConfiguration represents the UserGroupResolutionConfiguration schema from the OpenAPI specification
type UserGroupResolutionConfiguration struct {
	Usergroupresolutionmode interface{} `json:"UserGroupResolutionMode"`
}

// AttributeFilter represents the AttributeFilter schema from the OpenAPI specification
type AttributeFilter struct {
	Greaterthanorequals interface{} `json:"GreaterThanOrEquals,omitempty"`
	Containsall interface{} `json:"ContainsAll,omitempty"`
	Notfilter interface{} `json:"NotFilter,omitempty"`
	Equalsto interface{} `json:"EqualsTo,omitempty"`
	Greaterthan interface{} `json:"GreaterThan,omitempty"`
	Lessthan interface{} `json:"LessThan,omitempty"`
	Lessthanorequals interface{} `json:"LessThanOrEquals,omitempty"`
	Andallfilters interface{} `json:"AndAllFilters,omitempty"`
	Orallfilters interface{} `json:"OrAllFilters,omitempty"`
	Containsany interface{} `json:"ContainsAny,omitempty"`
}

// Highlight represents the Highlight schema from the OpenAPI specification
type Highlight struct {
	Topanswer interface{} `json:"TopAnswer,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Beginoffset interface{} `json:"BeginOffset"`
	Endoffset interface{} `json:"EndOffset"`
}

// S3Path represents the S3Path schema from the OpenAPI specification
type S3Path struct {
	Key interface{} `json:"Key"`
	Bucket interface{} `json:"Bucket"`
}

// CreateFeaturedResultsSetResponse represents the CreateFeaturedResultsSetResponse schema from the OpenAPI specification
type CreateFeaturedResultsSetResponse struct {
	Featuredresultsset interface{} `json:"FeaturedResultsSet,omitempty"`
}

// DescribePrincipalMappingRequest represents the DescribePrincipalMappingRequest schema from the OpenAPI specification
type DescribePrincipalMappingRequest struct {
	Datasourceid interface{} `json:"DataSourceId,omitempty"`
	Groupid interface{} `json:"GroupId"`
	Indexid interface{} `json:"IndexId"`
}

// FeaturedDocument represents the FeaturedDocument schema from the OpenAPI specification
type FeaturedDocument struct {
	Id interface{} `json:"Id,omitempty"`
}

// DocumentMetadataConfiguration represents the DocumentMetadataConfiguration schema from the OpenAPI specification
type DocumentMetadataConfiguration struct {
	TypeField interface{} `json:"Type"`
	Name interface{} `json:"Name"`
	Relevance interface{} `json:"Relevance,omitempty"`
	Search interface{} `json:"Search,omitempty"`
}

// DescribeFeaturedResultsSetRequest represents the DescribeFeaturedResultsSetRequest schema from the OpenAPI specification
type DescribeFeaturedResultsSetRequest struct {
	Featuredresultssetid interface{} `json:"FeaturedResultsSetId"`
	Indexid interface{} `json:"IndexId"`
}

// ListQuerySuggestionsBlockListsResponse represents the ListQuerySuggestionsBlockListsResponse schema from the OpenAPI specification
type ListQuerySuggestionsBlockListsResponse struct {
	Blocklistsummaryitems interface{} `json:"BlockListSummaryItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// BatchDeleteDocumentResponse represents the BatchDeleteDocumentResponse schema from the OpenAPI specification
type BatchDeleteDocumentResponse struct {
	Faileddocuments interface{} `json:"FailedDocuments,omitempty"`
}

// DataSourceSyncJobMetricTarget represents the DataSourceSyncJobMetricTarget schema from the OpenAPI specification
type DataSourceSyncJobMetricTarget struct {
	Datasourceid interface{} `json:"DataSourceId"`
	Datasourcesyncjobid interface{} `json:"DataSourceSyncJobId,omitempty"`
}

// AccessControlConfigurationSummary represents the AccessControlConfigurationSummary schema from the OpenAPI specification
type AccessControlConfigurationSummary struct {
	Id interface{} `json:"Id"`
}

// RelevanceFeedback represents the RelevanceFeedback schema from the OpenAPI specification
type RelevanceFeedback struct {
	Resultid interface{} `json:"ResultId"`
	Relevancevalue interface{} `json:"RelevanceValue"`
}

// AdditionalResultAttribute represents the AdditionalResultAttribute schema from the OpenAPI specification
type AdditionalResultAttribute struct {
	Valuetype interface{} `json:"ValueType"`
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// ProxyConfiguration represents the ProxyConfiguration schema from the OpenAPI specification
type ProxyConfiguration struct {
	Port interface{} `json:"Port"`
	Credentials interface{} `json:"Credentials,omitempty"`
	Host interface{} `json:"Host"`
}

// ExperiencesSummary represents the ExperiencesSummary schema from the OpenAPI specification
type ExperiencesSummary struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// ListEntityPersonasResponse represents the ListEntityPersonasResponse schema from the OpenAPI specification
type ListEntityPersonasResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Summaryitems interface{} `json:"SummaryItems,omitempty"`
}

// CreateIndexRequest represents the CreateIndexRequest schema from the OpenAPI specification
type CreateIndexRequest struct {
	Serversideencryptionconfiguration interface{} `json:"ServerSideEncryptionConfiguration,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Rolearn interface{} `json:"RoleArn"`
	Usercontextpolicy interface{} `json:"UserContextPolicy,omitempty"`
	Usertokenconfigurations interface{} `json:"UserTokenConfigurations,omitempty"`
	Edition interface{} `json:"Edition,omitempty"`
	Usergroupresolutionconfiguration interface{} `json:"UserGroupResolutionConfiguration,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
}

// ConfluenceSpaceToIndexFieldMapping represents the ConfluenceSpaceToIndexFieldMapping schema from the OpenAPI specification
type ConfluenceSpaceToIndexFieldMapping struct {
	Indexfieldname interface{} `json:"IndexFieldName,omitempty"`
	Datasourcefieldname interface{} `json:"DataSourceFieldName,omitempty"`
	Datefieldformat interface{} `json:"DateFieldFormat,omitempty"`
}

// BatchDeleteFeaturedResultsSetResponse represents the BatchDeleteFeaturedResultsSetResponse schema from the OpenAPI specification
type BatchDeleteFeaturedResultsSetResponse struct {
	Errors interface{} `json:"Errors"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// ClearQuerySuggestionsRequest represents the ClearQuerySuggestionsRequest schema from the OpenAPI specification
type ClearQuerySuggestionsRequest struct {
	Indexid interface{} `json:"IndexId"`
}

// QueryResultItem represents the QueryResultItem schema from the OpenAPI specification
type QueryResultItem struct {
	Documentattributes interface{} `json:"DocumentAttributes,omitempty"`
	Documentid interface{} `json:"DocumentId,omitempty"`
	Feedbacktoken interface{} `json:"FeedbackToken,omitempty"`
	Documenttitle interface{} `json:"DocumentTitle,omitempty"`
	Format interface{} `json:"Format,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Tableexcerpt interface{} `json:"TableExcerpt,omitempty"`
	Documenturi interface{} `json:"DocumentURI,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Documentexcerpt interface{} `json:"DocumentExcerpt,omitempty"`
	Scoreattributes interface{} `json:"ScoreAttributes,omitempty"`
	Additionalattributes interface{} `json:"AdditionalAttributes,omitempty"`
}

// UpdateAccessControlConfigurationRequest represents the UpdateAccessControlConfigurationRequest schema from the OpenAPI specification
type UpdateAccessControlConfigurationRequest struct {
	Accesscontrollist interface{} `json:"AccessControlList,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Hierarchicalaccesscontrollist interface{} `json:"HierarchicalAccessControlList,omitempty"`
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name,omitempty"`
}

// SalesforceStandardObjectConfiguration represents the SalesforceStandardObjectConfiguration schema from the OpenAPI specification
type SalesforceStandardObjectConfiguration struct {
	Name interface{} `json:"Name"`
	Documentdatafieldname interface{} `json:"DocumentDataFieldName"`
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
}

// DeleteThesaurusRequest represents the DeleteThesaurusRequest schema from the OpenAPI specification
type DeleteThesaurusRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// DescribeDataSourceResponse represents the DescribeDataSourceResponse schema from the OpenAPI specification
type DescribeDataSourceResponse struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Customdocumentenrichmentconfiguration interface{} `json:"CustomDocumentEnrichmentConfiguration,omitempty"`
	Indexid interface{} `json:"IndexId,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
}

// ListQuerySuggestionsBlockListsRequest represents the ListQuerySuggestionsBlockListsRequest schema from the OpenAPI specification
type ListQuerySuggestionsBlockListsRequest struct {
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Urls represents the Urls schema from the OpenAPI specification
type Urls struct {
	Seedurlconfiguration interface{} `json:"SeedUrlConfiguration,omitempty"`
	Sitemapsconfiguration interface{} `json:"SiteMapsConfiguration,omitempty"`
}

// ListGroupsOlderThanOrderingIdResponse represents the ListGroupsOlderThanOrderingIdResponse schema from the OpenAPI specification
type ListGroupsOlderThanOrderingIdResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Groupssummaries interface{} `json:"GroupsSummaries,omitempty"`
}

// DeleteDataSourceRequest represents the DeleteDataSourceRequest schema from the OpenAPI specification
type DeleteDataSourceRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// DeleteFaqRequest represents the DeleteFaqRequest schema from the OpenAPI specification
type DeleteFaqRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// AssociatePersonasToEntitiesRequest represents the AssociatePersonasToEntitiesRequest schema from the OpenAPI specification
type AssociatePersonasToEntitiesRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Personas interface{} `json:"Personas"`
}

// UserTokenConfiguration represents the UserTokenConfiguration schema from the OpenAPI specification
type UserTokenConfiguration struct {
	Jsontokentypeconfiguration interface{} `json:"JsonTokenTypeConfiguration,omitempty"`
	Jwttokentypeconfiguration interface{} `json:"JwtTokenTypeConfiguration,omitempty"`
}

// GetSnapshotsResponse represents the GetSnapshotsResponse schema from the OpenAPI specification
type GetSnapshotsResponse struct {
	Snapshottimefilter interface{} `json:"SnapShotTimeFilter,omitempty"`
	Snapshotsdata interface{} `json:"SnapshotsData,omitempty"`
	Snapshotsdataheader interface{} `json:"SnapshotsDataHeader,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// OneDriveUsers represents the OneDriveUsers schema from the OpenAPI specification
type OneDriveUsers struct {
	Onedriveuserlist interface{} `json:"OneDriveUserList,omitempty"`
	Onedriveusers3path interface{} `json:"OneDriveUserS3Path,omitempty"`
}

// DescribeQuerySuggestionsBlockListResponse represents the DescribeQuerySuggestionsBlockListResponse schema from the OpenAPI specification
type DescribeQuerySuggestionsBlockListResponse struct {
	Itemcount interface{} `json:"ItemCount,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Indexid interface{} `json:"IndexId,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Sources3path interface{} `json:"SourceS3Path,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Filesizebytes interface{} `json:"FileSizeBytes,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
}

// ListFeaturedResultsSetsResponse represents the ListFeaturedResultsSetsResponse schema from the OpenAPI specification
type ListFeaturedResultsSetsResponse struct {
	Featuredresultssetsummaryitems interface{} `json:"FeaturedResultsSetSummaryItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GoogleDriveConfiguration represents the GoogleDriveConfiguration schema from the OpenAPI specification
type GoogleDriveConfiguration struct {
	Secretarn interface{} `json:"SecretArn"`
	Excludemimetypes interface{} `json:"ExcludeMimeTypes,omitempty"`
	Excludeshareddrives interface{} `json:"ExcludeSharedDrives,omitempty"`
	Excludeuseraccounts interface{} `json:"ExcludeUserAccounts,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
}

// DocumentAttributeCondition represents the DocumentAttributeCondition schema from the OpenAPI specification
type DocumentAttributeCondition struct {
	Conditiondocumentattributekey interface{} `json:"ConditionDocumentAttributeKey"`
	Conditiononvalue interface{} `json:"ConditionOnValue,omitempty"`
	Operator interface{} `json:"Operator"`
}

// DescribeQuerySuggestionsConfigResponse represents the DescribeQuerySuggestionsConfigResponse schema from the OpenAPI specification
type DescribeQuerySuggestionsConfigResponse struct {
	Minimumquerycount interface{} `json:"MinimumQueryCount,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Queryloglookbackwindowindays interface{} `json:"QueryLogLookBackWindowInDays,omitempty"`
	Totalsuggestionscount interface{} `json:"TotalSuggestionsCount,omitempty"`
	Attributesuggestionsconfig interface{} `json:"AttributeSuggestionsConfig,omitempty"`
	Includequerieswithoutuserinformation interface{} `json:"IncludeQueriesWithoutUserInformation,omitempty"`
	Lastcleartime interface{} `json:"LastClearTime,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Lastsuggestionsbuildtime interface{} `json:"LastSuggestionsBuildTime,omitempty"`
	Minimumnumberofqueryingusers interface{} `json:"MinimumNumberOfQueryingUsers,omitempty"`
}

// ListFaqsResponse represents the ListFaqsResponse schema from the OpenAPI specification
type ListFaqsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Faqsummaryitems interface{} `json:"FaqSummaryItems,omitempty"`
}

// GroupSummary represents the GroupSummary schema from the OpenAPI specification
type GroupSummary struct {
	Groupid interface{} `json:"GroupId,omitempty"`
	Orderingid interface{} `json:"OrderingId,omitempty"`
}

// BoxConfiguration represents the BoxConfiguration schema from the OpenAPI specification
type BoxConfiguration struct {
	Weblinkfieldmappings interface{} `json:"WebLinkFieldMappings,omitempty"`
	Commentfieldmappings interface{} `json:"CommentFieldMappings,omitempty"`
	Crawltasks interface{} `json:"CrawlTasks,omitempty"`
	Crawlweblinks interface{} `json:"CrawlWebLinks,omitempty"`
	Enterpriseid interface{} `json:"EnterpriseId"`
	Filefieldmappings interface{} `json:"FileFieldMappings,omitempty"`
	Crawlcomments interface{} `json:"CrawlComments,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Taskfieldmappings interface{} `json:"TaskFieldMappings,omitempty"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Usechangelog interface{} `json:"UseChangeLog,omitempty"`
}

// IndexConfigurationSummary represents the IndexConfigurationSummary schema from the OpenAPI specification
type IndexConfigurationSummary struct {
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status"`
	Updatedat interface{} `json:"UpdatedAt"`
	Createdat interface{} `json:"CreatedAt"`
	Edition interface{} `json:"Edition,omitempty"`
}

// ListDataSourcesRequest represents the ListDataSourcesRequest schema from the OpenAPI specification
type ListDataSourcesRequest struct {
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DisassociateEntitiesFromExperienceRequest represents the DisassociateEntitiesFromExperienceRequest schema from the OpenAPI specification
type DisassociateEntitiesFromExperienceRequest struct {
	Entitylist interface{} `json:"EntityList"`
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// SuggestionHighlight represents the SuggestionHighlight schema from the OpenAPI specification
type SuggestionHighlight struct {
	Beginoffset interface{} `json:"BeginOffset,omitempty"`
	Endoffset interface{} `json:"EndOffset,omitempty"`
}

// ExperienceEndpoint represents the ExperienceEndpoint schema from the OpenAPI specification
type ExperienceEndpoint struct {
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Endpointtype interface{} `json:"EndpointType,omitempty"`
}

// DisassociateEntitiesFromExperienceResponse represents the DisassociateEntitiesFromExperienceResponse schema from the OpenAPI specification
type DisassociateEntitiesFromExperienceResponse struct {
	Failedentitylist interface{} `json:"FailedEntityList,omitempty"`
}

// ClickFeedback represents the ClickFeedback schema from the OpenAPI specification
type ClickFeedback struct {
	Clicktime interface{} `json:"ClickTime"`
	Resultid interface{} `json:"ResultId"`
}

// FaqSummary represents the FaqSummary schema from the OpenAPI specification
type FaqSummary struct {
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Fileformat interface{} `json:"FileFormat,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// ListThesauriRequest represents the ListThesauriRequest schema from the OpenAPI specification
type ListThesauriRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ConfluenceAttachmentToIndexFieldMapping represents the ConfluenceAttachmentToIndexFieldMapping schema from the OpenAPI specification
type ConfluenceAttachmentToIndexFieldMapping struct {
	Datefieldformat interface{} `json:"DateFieldFormat,omitempty"`
	Indexfieldname interface{} `json:"IndexFieldName,omitempty"`
	Datasourcefieldname interface{} `json:"DataSourceFieldName,omitempty"`
}

// DescribeThesaurusResponse represents the DescribeThesaurusResponse schema from the OpenAPI specification
type DescribeThesaurusResponse struct {
	Name interface{} `json:"Name,omitempty"`
	Synonymrulecount interface{} `json:"SynonymRuleCount,omitempty"`
	Termcount interface{} `json:"TermCount,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Filesizebytes interface{} `json:"FileSizeBytes,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Indexid interface{} `json:"IndexId,omitempty"`
	Sources3path S3Path `json:"SourceS3Path,omitempty"` // Information required to find a specific file in an Amazon S3 bucket.
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
}

// TemplateConfiguration represents the TemplateConfiguration schema from the OpenAPI specification
type TemplateConfiguration struct {
	Template interface{} `json:"Template,omitempty"`
}

// QueryResult represents the QueryResult schema from the OpenAPI specification
type QueryResult struct {
	Totalnumberofresults interface{} `json:"TotalNumberOfResults,omitempty"`
	Warnings interface{} `json:"Warnings,omitempty"`
	Facetresults interface{} `json:"FacetResults,omitempty"`
	Featuredresultsitems interface{} `json:"FeaturedResultsItems,omitempty"`
	Queryid interface{} `json:"QueryId,omitempty"`
	Resultitems interface{} `json:"ResultItems,omitempty"`
	Spellcorrectedqueries interface{} `json:"SpellCorrectedQueries,omitempty"`
}

// SalesforceStandardObjectAttachmentConfiguration represents the SalesforceStandardObjectAttachmentConfiguration schema from the OpenAPI specification
type SalesforceStandardObjectAttachmentConfiguration struct {
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
}

// CreateDataSourceRequest represents the CreateDataSourceRequest schema from the OpenAPI specification
type CreateDataSourceRequest struct {
	TypeField interface{} `json:"Type"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Customdocumentenrichmentconfiguration interface{} `json:"CustomDocumentEnrichmentConfiguration,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
}

// SharePointConfiguration represents the SharePointConfiguration schema from the OpenAPI specification
type SharePointConfiguration struct {
	Usechangelog interface{} `json:"UseChangeLog,omitempty"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Proxyconfiguration interface{} `json:"ProxyConfiguration,omitempty"`
	Sslcertificates3path interface{} `json:"SslCertificateS3Path,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Authenticationtype interface{} `json:"AuthenticationType,omitempty"`
	Crawlattachments interface{} `json:"CrawlAttachments,omitempty"`
	Disablelocalgroups interface{} `json:"DisableLocalGroups,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Urls interface{} `json:"Urls"`
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Sharepointversion interface{} `json:"SharePointVersion"`
}

// DescribeFeaturedResultsSetResponse represents the DescribeFeaturedResultsSetResponse schema from the OpenAPI specification
type DescribeFeaturedResultsSetResponse struct {
	Featureddocumentswithmetadata interface{} `json:"FeaturedDocumentsWithMetadata,omitempty"`
	Featuredresultssetid interface{} `json:"FeaturedResultsSetId,omitempty"`
	Featuredresultssetname interface{} `json:"FeaturedResultsSetName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Lastupdatedtimestamp interface{} `json:"LastUpdatedTimestamp,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Creationtimestamp interface{} `json:"CreationTimestamp,omitempty"`
	Featureddocumentsmissing interface{} `json:"FeaturedDocumentsMissing,omitempty"`
	Querytexts interface{} `json:"QueryTexts,omitempty"`
}

// CreateFaqRequest represents the CreateFaqRequest schema from the OpenAPI specification
type CreateFaqRequest struct {
	Indexid interface{} `json:"IndexId"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Fileformat interface{} `json:"FileFormat,omitempty"`
	S3path interface{} `json:"S3Path"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// QuipConfiguration represents the QuipConfiguration schema from the OpenAPI specification
type QuipConfiguration struct {
	Domain interface{} `json:"Domain"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Folderids interface{} `json:"FolderIds,omitempty"`
	Messagefieldmappings interface{} `json:"MessageFieldMappings,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Crawlfilecomments interface{} `json:"CrawlFileComments,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Threadfieldmappings interface{} `json:"ThreadFieldMappings,omitempty"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Attachmentfieldmappings interface{} `json:"AttachmentFieldMappings,omitempty"`
	Crawlchatrooms interface{} `json:"CrawlChatRooms,omitempty"`
	Crawlattachments interface{} `json:"CrawlAttachments,omitempty"`
}

// FeaturedResultsSet represents the FeaturedResultsSet schema from the OpenAPI specification
type FeaturedResultsSet struct {
	Status interface{} `json:"Status,omitempty"`
	Creationtimestamp interface{} `json:"CreationTimestamp,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Featureddocuments interface{} `json:"FeaturedDocuments,omitempty"`
	Featuredresultssetid interface{} `json:"FeaturedResultsSetId,omitempty"`
	Featuredresultssetname interface{} `json:"FeaturedResultsSetName,omitempty"`
	Lastupdatedtimestamp interface{} `json:"LastUpdatedTimestamp,omitempty"`
	Querytexts interface{} `json:"QueryTexts,omitempty"`
}

// BatchPutDocumentResponseFailedDocument represents the BatchPutDocumentResponseFailedDocument schema from the OpenAPI specification
type BatchPutDocumentResponseFailedDocument struct {
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
}

// RetrieveResultItem represents the RetrieveResultItem schema from the OpenAPI specification
type RetrieveResultItem struct {
	Documenturi interface{} `json:"DocumentURI,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Content interface{} `json:"Content,omitempty"`
	Documentattributes interface{} `json:"DocumentAttributes,omitempty"`
	Documentid interface{} `json:"DocumentId,omitempty"`
	Documenttitle interface{} `json:"DocumentTitle,omitempty"`
}

// ListFeaturedResultsSetsRequest represents the ListFeaturedResultsSetsRequest schema from the OpenAPI specification
type ListFeaturedResultsSetsRequest struct {
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// TextWithHighlights represents the TextWithHighlights schema from the OpenAPI specification
type TextWithHighlights struct {
	Highlights interface{} `json:"Highlights,omitempty"`
	Text interface{} `json:"Text,omitempty"`
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
	Title interface{} `json:"Title,omitempty"`
	Accesscontrolconfigurationid interface{} `json:"AccessControlConfigurationId,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Hierarchicalaccesscontrollist interface{} `json:"HierarchicalAccessControlList,omitempty"`
	S3path S3Path `json:"S3Path,omitempty"` // Information required to find a specific file in an Amazon S3 bucket.
	Contenttype interface{} `json:"ContentType,omitempty"`
	Accesscontrollist interface{} `json:"AccessControlList,omitempty"`
	Blob interface{} `json:"Blob,omitempty"`
	Id interface{} `json:"Id"`
}

// ColumnConfiguration represents the ColumnConfiguration schema from the OpenAPI specification
type ColumnConfiguration struct {
	Changedetectingcolumns interface{} `json:"ChangeDetectingColumns"`
	Documentdatacolumnname interface{} `json:"DocumentDataColumnName"`
	Documentidcolumnname interface{} `json:"DocumentIdColumnName"`
	Documenttitlecolumnname interface{} `json:"DocumentTitleColumnName,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
}

// TableExcerpt represents the TableExcerpt schema from the OpenAPI specification
type TableExcerpt struct {
	Rows interface{} `json:"Rows,omitempty"`
	Totalnumberofrows interface{} `json:"TotalNumberOfRows,omitempty"`
}

// SpellCorrectedQuery represents the SpellCorrectedQuery schema from the OpenAPI specification
type SpellCorrectedQuery struct {
	Corrections interface{} `json:"Corrections,omitempty"`
	Suggestedquerytext interface{} `json:"SuggestedQueryText,omitempty"`
}

// ConfluenceAttachmentConfiguration represents the ConfluenceAttachmentConfiguration schema from the OpenAPI specification
type ConfluenceAttachmentConfiguration struct {
	Attachmentfieldmappings interface{} `json:"AttachmentFieldMappings,omitempty"`
	Crawlattachments interface{} `json:"CrawlAttachments,omitempty"`
}

// DisassociatePersonasFromEntitiesRequest represents the DisassociatePersonasFromEntitiesRequest schema from the OpenAPI specification
type DisassociatePersonasFromEntitiesRequest struct {
	Entityids interface{} `json:"EntityIds"`
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// ServiceNowServiceCatalogConfiguration represents the ServiceNowServiceCatalogConfiguration schema from the OpenAPI specification
type ServiceNowServiceCatalogConfiguration struct {
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
	Excludeattachmentfilepatterns interface{} `json:"ExcludeAttachmentFilePatterns,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Includeattachmentfilepatterns interface{} `json:"IncludeAttachmentFilePatterns,omitempty"`
	Crawlattachments interface{} `json:"CrawlAttachments,omitempty"`
	Documentdatafieldname interface{} `json:"DocumentDataFieldName"`
}

// CreateAccessControlConfigurationRequest represents the CreateAccessControlConfigurationRequest schema from the OpenAPI specification
type CreateAccessControlConfigurationRequest struct {
	Accesscontrollist interface{} `json:"AccessControlList,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Hierarchicalaccesscontrollist interface{} `json:"HierarchicalAccessControlList,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name"`
}

// AssociateEntitiesToExperienceRequest represents the AssociateEntitiesToExperienceRequest schema from the OpenAPI specification
type AssociateEntitiesToExperienceRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Entitylist interface{} `json:"EntityList"`
}

// FeaturedResultsSetSummary represents the FeaturedResultsSetSummary schema from the OpenAPI specification
type FeaturedResultsSetSummary struct {
	Status interface{} `json:"Status,omitempty"`
	Creationtimestamp interface{} `json:"CreationTimestamp,omitempty"`
	Featuredresultssetid interface{} `json:"FeaturedResultsSetId,omitempty"`
	Featuredresultssetname interface{} `json:"FeaturedResultsSetName,omitempty"`
	Lastupdatedtimestamp interface{} `json:"LastUpdatedTimestamp,omitempty"`
}

// ConfluenceBlogConfiguration represents the ConfluenceBlogConfiguration schema from the OpenAPI specification
type ConfluenceBlogConfiguration struct {
	Blogfieldmappings interface{} `json:"BlogFieldMappings,omitempty"`
}

// PutPrincipalMappingRequest represents the PutPrincipalMappingRequest schema from the OpenAPI specification
type PutPrincipalMappingRequest struct {
	Datasourceid interface{} `json:"DataSourceId,omitempty"`
	Groupid interface{} `json:"GroupId"`
	Groupmembers interface{} `json:"GroupMembers"`
	Indexid interface{} `json:"IndexId"`
	Orderingid interface{} `json:"OrderingId,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
}

// ConfluencePageToIndexFieldMapping represents the ConfluencePageToIndexFieldMapping schema from the OpenAPI specification
type ConfluencePageToIndexFieldMapping struct {
	Datasourcefieldname interface{} `json:"DataSourceFieldName,omitempty"`
	Datefieldformat interface{} `json:"DateFieldFormat,omitempty"`
	Indexfieldname interface{} `json:"IndexFieldName,omitempty"`
}

// FailedEntity represents the FailedEntity schema from the OpenAPI specification
type FailedEntity struct {
	Entityid interface{} `json:"EntityId,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// BatchDeleteFeaturedResultsSetRequest represents the BatchDeleteFeaturedResultsSetRequest schema from the OpenAPI specification
type BatchDeleteFeaturedResultsSetRequest struct {
	Featuredresultssetids interface{} `json:"FeaturedResultsSetIds"`
	Indexid interface{} `json:"IndexId"`
}

// Principal represents the Principal schema from the OpenAPI specification
type Principal struct {
	Name interface{} `json:"Name"`
	TypeField interface{} `json:"Type"`
	Access interface{} `json:"Access"`
	Datasourceid interface{} `json:"DataSourceId,omitempty"`
}

// AuthenticationConfiguration represents the AuthenticationConfiguration schema from the OpenAPI specification
type AuthenticationConfiguration struct {
	Basicauthentication interface{} `json:"BasicAuthentication,omitempty"`
}

// ConfluenceBlogToIndexFieldMapping represents the ConfluenceBlogToIndexFieldMapping schema from the OpenAPI specification
type ConfluenceBlogToIndexFieldMapping struct {
	Datefieldformat interface{} `json:"DateFieldFormat,omitempty"`
	Indexfieldname interface{} `json:"IndexFieldName,omitempty"`
	Datasourcefieldname interface{} `json:"DataSourceFieldName,omitempty"`
}

// GitHubDocumentCrawlProperties represents the GitHubDocumentCrawlProperties schema from the OpenAPI specification
type GitHubDocumentCrawlProperties struct {
	Crawlissuecommentattachment interface{} `json:"CrawlIssueCommentAttachment,omitempty"`
	Crawlpullrequest interface{} `json:"CrawlPullRequest,omitempty"`
	Crawlpullrequestcomment interface{} `json:"CrawlPullRequestComment,omitempty"`
	Crawlpullrequestcommentattachment interface{} `json:"CrawlPullRequestCommentAttachment,omitempty"`
	Crawlrepositorydocuments interface{} `json:"CrawlRepositoryDocuments,omitempty"`
	Crawlissue interface{} `json:"CrawlIssue,omitempty"`
	Crawlissuecomment interface{} `json:"CrawlIssueComment,omitempty"`
}

// StartDataSourceSyncJobRequest represents the StartDataSourceSyncJobRequest schema from the OpenAPI specification
type StartDataSourceSyncJobRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// ListGroupsOlderThanOrderingIdRequest represents the ListGroupsOlderThanOrderingIdRequest schema from the OpenAPI specification
type ListGroupsOlderThanOrderingIdRequest struct {
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Orderingid interface{} `json:"OrderingId"`
	Datasourceid interface{} `json:"DataSourceId,omitempty"`
}

// ServiceNowKnowledgeArticleConfiguration represents the ServiceNowKnowledgeArticleConfiguration schema from the OpenAPI specification
type ServiceNowKnowledgeArticleConfiguration struct {
	Crawlattachments interface{} `json:"CrawlAttachments,omitempty"`
	Documentdatafieldname interface{} `json:"DocumentDataFieldName"`
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
	Excludeattachmentfilepatterns interface{} `json:"ExcludeAttachmentFilePatterns,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Filterquery interface{} `json:"FilterQuery,omitempty"`
	Includeattachmentfilepatterns interface{} `json:"IncludeAttachmentFilePatterns,omitempty"`
}

// ThesaurusSummary represents the ThesaurusSummary schema from the OpenAPI specification
type ThesaurusSummary struct {
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// ScoreAttributes represents the ScoreAttributes schema from the OpenAPI specification
type ScoreAttributes struct {
	Scoreconfidence interface{} `json:"ScoreConfidence,omitempty"`
}

// CreateFeaturedResultsSetRequest represents the CreateFeaturedResultsSetRequest schema from the OpenAPI specification
type CreateFeaturedResultsSetRequest struct {
	Featureddocuments interface{} `json:"FeaturedDocuments,omitempty"`
	Featuredresultssetname interface{} `json:"FeaturedResultsSetName"`
	Indexid interface{} `json:"IndexId"`
	Querytexts interface{} `json:"QueryTexts,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// ListEntityPersonasRequest represents the ListEntityPersonasRequest schema from the OpenAPI specification
type ListEntityPersonasRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MemberGroup represents the MemberGroup schema from the OpenAPI specification
type MemberGroup struct {
	Datasourceid interface{} `json:"DataSourceId,omitempty"`
	Groupid interface{} `json:"GroupId"`
}

// DescribeIndexRequest represents the DescribeIndexRequest schema from the OpenAPI specification
type DescribeIndexRequest struct {
	Id interface{} `json:"Id"`
}

// GetQuerySuggestionsRequest represents the GetQuerySuggestionsRequest schema from the OpenAPI specification
type GetQuerySuggestionsRequest struct {
	Attributesuggestionsconfig interface{} `json:"AttributeSuggestionsConfig,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Maxsuggestionscount interface{} `json:"MaxSuggestionsCount,omitempty"`
	Querytext interface{} `json:"QueryText"`
	Suggestiontypes interface{} `json:"SuggestionTypes,omitempty"`
}

// SuggestableConfig represents the SuggestableConfig schema from the OpenAPI specification
type SuggestableConfig struct {
	Attributename interface{} `json:"AttributeName,omitempty"`
	Suggestable interface{} `json:"Suggestable,omitempty"`
}

// ExperienceEntitiesSummary represents the ExperienceEntitiesSummary schema from the OpenAPI specification
type ExperienceEntitiesSummary struct {
	Displaydata interface{} `json:"DisplayData,omitempty"`
	Entityid interface{} `json:"EntityId,omitempty"`
	Entitytype interface{} `json:"EntityType,omitempty"`
}

// SalesforceConfiguration represents the SalesforceConfiguration schema from the OpenAPI specification
type SalesforceConfiguration struct {
	Chatterfeedconfiguration interface{} `json:"ChatterFeedConfiguration,omitempty"`
	Excludeattachmentfilepatterns interface{} `json:"ExcludeAttachmentFilePatterns,omitempty"`
	Includeattachmentfilepatterns interface{} `json:"IncludeAttachmentFilePatterns,omitempty"`
	Knowledgearticleconfiguration interface{} `json:"KnowledgeArticleConfiguration,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Standardobjectattachmentconfiguration interface{} `json:"StandardObjectAttachmentConfiguration,omitempty"`
	Crawlattachments interface{} `json:"CrawlAttachments,omitempty"`
	Serverurl interface{} `json:"ServerUrl"`
	Standardobjectconfigurations interface{} `json:"StandardObjectConfigurations,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// DataSourceToIndexFieldMapping represents the DataSourceToIndexFieldMapping schema from the OpenAPI specification
type DataSourceToIndexFieldMapping struct {
	Indexfieldname interface{} `json:"IndexFieldName"`
	Datasourcefieldname interface{} `json:"DataSourceFieldName"`
	Datefieldformat interface{} `json:"DateFieldFormat,omitempty"`
}

// DataSourceVpcConfiguration represents the DataSourceVpcConfiguration schema from the OpenAPI specification
type DataSourceVpcConfiguration struct {
	Securitygroupids interface{} `json:"SecurityGroupIds"`
	Subnetids interface{} `json:"SubnetIds"`
}

// DescribeFaqResponse represents the DescribeFaqResponse schema from the OpenAPI specification
type DescribeFaqResponse struct {
	Indexid interface{} `json:"IndexId,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	S3path S3Path `json:"S3Path,omitempty"` // Information required to find a specific file in an Amazon S3 bucket.
	Status interface{} `json:"Status,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Fileformat interface{} `json:"FileFormat,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// UpdateFeaturedResultsSetRequest represents the UpdateFeaturedResultsSetRequest schema from the OpenAPI specification
type UpdateFeaturedResultsSetRequest struct {
	Featureddocuments interface{} `json:"FeaturedDocuments,omitempty"`
	Featuredresultssetid interface{} `json:"FeaturedResultsSetId"`
	Featuredresultssetname interface{} `json:"FeaturedResultsSetName,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Querytexts interface{} `json:"QueryTexts,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// GroupOrderingIdSummary represents the GroupOrderingIdSummary schema from the OpenAPI specification
type GroupOrderingIdSummary struct {
	Receivedat interface{} `json:"ReceivedAt,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastupdatedat interface{} `json:"LastUpdatedAt,omitempty"`
	Orderingid interface{} `json:"OrderingId,omitempty"`
}

// QuerySuggestionsBlockListSummary represents the QuerySuggestionsBlockListSummary schema from the OpenAPI specification
type QuerySuggestionsBlockListSummary struct {
	Id interface{} `json:"Id,omitempty"`
	Itemcount interface{} `json:"ItemCount,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// DeleteAccessControlConfigurationRequest represents the DeleteAccessControlConfigurationRequest schema from the OpenAPI specification
type DeleteAccessControlConfigurationRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// SlackConfiguration represents the SlackConfiguration schema from the OpenAPI specification
type SlackConfiguration struct {
	Lookbackperiod interface{} `json:"LookBackPeriod,omitempty"`
	Privatechannelfilter interface{} `json:"PrivateChannelFilter,omitempty"`
	Publicchannelfilter interface{} `json:"PublicChannelFilter,omitempty"`
	Slackentitylist interface{} `json:"SlackEntityList"`
	Usechangelog interface{} `json:"UseChangeLog,omitempty"`
	Crawlbotmessage interface{} `json:"CrawlBotMessage,omitempty"`
	Excludearchived interface{} `json:"ExcludeArchived,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Sincecrawldate interface{} `json:"SinceCrawlDate"`
	Teamid interface{} `json:"TeamId"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
}

// ContentSourceConfiguration represents the ContentSourceConfiguration schema from the OpenAPI specification
type ContentSourceConfiguration struct {
	Datasourceids interface{} `json:"DataSourceIds,omitempty"`
	Directputcontent interface{} `json:"DirectPutContent,omitempty"`
	Faqids interface{} `json:"FaqIds,omitempty"`
}

// CreateExperienceResponse represents the CreateExperienceResponse schema from the OpenAPI specification
type CreateExperienceResponse struct {
	Id interface{} `json:"Id"`
}

// UserContext represents the UserContext schema from the OpenAPI specification
type UserContext struct {
	Datasourcegroups interface{} `json:"DataSourceGroups,omitempty"`
	Groups interface{} `json:"Groups,omitempty"`
	Token interface{} `json:"Token,omitempty"`
	Userid interface{} `json:"UserId,omitempty"`
}

// GetQuerySuggestionsResponse represents the GetQuerySuggestionsResponse schema from the OpenAPI specification
type GetQuerySuggestionsResponse struct {
	Querysuggestionsid interface{} `json:"QuerySuggestionsId,omitempty"`
	Suggestions interface{} `json:"Suggestions,omitempty"`
}

// FeaturedDocumentWithMetadata represents the FeaturedDocumentWithMetadata schema from the OpenAPI specification
type FeaturedDocumentWithMetadata struct {
	Id interface{} `json:"Id,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Uri interface{} `json:"URI,omitempty"`
}

// ListExperiencesRequest represents the ListExperiencesRequest schema from the OpenAPI specification
type ListExperiencesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Indexid interface{} `json:"IndexId"`
}

// TableCell represents the TableCell schema from the OpenAPI specification
type TableCell struct {
	Highlighted interface{} `json:"Highlighted,omitempty"`
	Topanswer interface{} `json:"TopAnswer,omitempty"`
	Value interface{} `json:"Value,omitempty"`
	Header interface{} `json:"Header,omitempty"`
}

// DescribeDataSourceRequest represents the DescribeDataSourceRequest schema from the OpenAPI specification
type DescribeDataSourceRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// DeleteIndexRequest represents the DeleteIndexRequest schema from the OpenAPI specification
type DeleteIndexRequest struct {
	Id interface{} `json:"Id"`
}

// CreateThesaurusRequest represents the CreateThesaurusRequest schema from the OpenAPI specification
type CreateThesaurusRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name"`
	Rolearn interface{} `json:"RoleArn"`
	Sources3path interface{} `json:"SourceS3Path"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
}

// FeaturedDocumentMissing represents the FeaturedDocumentMissing schema from the OpenAPI specification
type FeaturedDocumentMissing struct {
	Id interface{} `json:"Id,omitempty"`
}

// JiraConfiguration represents the JiraConfiguration schema from the OpenAPI specification
type JiraConfiguration struct {
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Jiraaccounturl interface{} `json:"JiraAccountUrl"`
	Secretarn interface{} `json:"SecretArn"`
	Worklogfieldmappings interface{} `json:"WorkLogFieldMappings,omitempty"`
	Issuesubentityfilter interface{} `json:"IssueSubEntityFilter,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Attachmentfieldmappings interface{} `json:"AttachmentFieldMappings,omitempty"`
	Commentfieldmappings interface{} `json:"CommentFieldMappings,omitempty"`
	Project interface{} `json:"Project,omitempty"`
	Issuefieldmappings interface{} `json:"IssueFieldMappings,omitempty"`
	Issuetype interface{} `json:"IssueType,omitempty"`
	Projectfieldmappings interface{} `json:"ProjectFieldMappings,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Usechangelog interface{} `json:"UseChangeLog,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
}

// ConfluenceConfiguration represents the ConfluenceConfiguration schema from the OpenAPI specification
type ConfluenceConfiguration struct {
	Blogconfiguration interface{} `json:"BlogConfiguration,omitempty"`
	Proxyconfiguration interface{} `json:"ProxyConfiguration,omitempty"`
	Attachmentconfiguration interface{} `json:"AttachmentConfiguration,omitempty"`
	Pageconfiguration interface{} `json:"PageConfiguration,omitempty"`
	Spaceconfiguration interface{} `json:"SpaceConfiguration,omitempty"`
	Version interface{} `json:"Version"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Authenticationtype interface{} `json:"AuthenticationType,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Serverurl interface{} `json:"ServerUrl"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
}

// JsonTokenTypeConfiguration represents the JsonTokenTypeConfiguration schema from the OpenAPI specification
type JsonTokenTypeConfiguration struct {
	Usernameattributefield interface{} `json:"UserNameAttributeField"`
	Groupattributefield interface{} `json:"GroupAttributeField"`
}

// DataSourceConfiguration represents the DataSourceConfiguration schema from the OpenAPI specification
type DataSourceConfiguration struct {
	Fsxconfiguration interface{} `json:"FsxConfiguration,omitempty"`
	Onedriveconfiguration interface{} `json:"OneDriveConfiguration,omitempty"`
	Boxconfiguration interface{} `json:"BoxConfiguration,omitempty"`
	Templateconfiguration interface{} `json:"TemplateConfiguration,omitempty"`
	Servicenowconfiguration interface{} `json:"ServiceNowConfiguration,omitempty"`
	Confluenceconfiguration interface{} `json:"ConfluenceConfiguration,omitempty"`
	Quipconfiguration interface{} `json:"QuipConfiguration,omitempty"`
	Sharepointconfiguration interface{} `json:"SharePointConfiguration,omitempty"`
	Alfrescoconfiguration interface{} `json:"AlfrescoConfiguration,omitempty"`
	Salesforceconfiguration interface{} `json:"SalesforceConfiguration,omitempty"`
	Slackconfiguration interface{} `json:"SlackConfiguration,omitempty"`
	Workdocsconfiguration interface{} `json:"WorkDocsConfiguration,omitempty"`
	Webcrawlerconfiguration WebCrawlerConfiguration `json:"WebCrawlerConfiguration,omitempty"` // Provides the configuration information required for Amazon Kendra Web Crawler.
	Githubconfiguration interface{} `json:"GitHubConfiguration,omitempty"`
	Googledriveconfiguration interface{} `json:"GoogleDriveConfiguration,omitempty"`
	S3configuration interface{} `json:"S3Configuration,omitempty"`
	Jiraconfiguration interface{} `json:"JiraConfiguration,omitempty"`
	Databaseconfiguration interface{} `json:"DatabaseConfiguration,omitempty"`
}

// WorkDocsConfiguration represents the WorkDocsConfiguration schema from the OpenAPI specification
type WorkDocsConfiguration struct {
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Organizationid interface{} `json:"OrganizationId"`
	Usechangelog interface{} `json:"UseChangeLog,omitempty"`
	Crawlcomments interface{} `json:"CrawlComments,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// DescribeAccessControlConfigurationResponse represents the DescribeAccessControlConfigurationResponse schema from the OpenAPI specification
type DescribeAccessControlConfigurationResponse struct {
	Hierarchicalaccesscontrollist interface{} `json:"HierarchicalAccessControlList,omitempty"`
	Name interface{} `json:"Name"`
	Accesscontrollist interface{} `json:"AccessControlList,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// UpdateQuerySuggestionsBlockListRequest represents the UpdateQuerySuggestionsBlockListRequest schema from the OpenAPI specification
type UpdateQuerySuggestionsBlockListRequest struct {
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Sources3path interface{} `json:"SourceS3Path,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name,omitempty"`
}

// CapacityUnitsConfiguration represents the CapacityUnitsConfiguration schema from the OpenAPI specification
type CapacityUnitsConfiguration struct {
	Querycapacityunits interface{} `json:"QueryCapacityUnits"`
	Storagecapacityunits interface{} `json:"StorageCapacityUnits"`
}

// SeedUrlConfiguration represents the SeedUrlConfiguration schema from the OpenAPI specification
type SeedUrlConfiguration struct {
	Seedurls interface{} `json:"SeedUrls"`
	Webcrawlermode interface{} `json:"WebCrawlerMode,omitempty"`
}

// DescribeAccessControlConfigurationRequest represents the DescribeAccessControlConfigurationRequest schema from the OpenAPI specification
type DescribeAccessControlConfigurationRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// FeaturedResultsItem represents the FeaturedResultsItem schema from the OpenAPI specification
type FeaturedResultsItem struct {
	Documentid interface{} `json:"DocumentId,omitempty"`
	Documenttitle TextWithHighlights `json:"DocumentTitle,omitempty"` // Provides text and information about where to highlight the text.
	Id interface{} `json:"Id,omitempty"`
	Documenturi interface{} `json:"DocumentURI,omitempty"`
	Feedbacktoken interface{} `json:"FeedbackToken,omitempty"`
	Additionalattributes interface{} `json:"AdditionalAttributes,omitempty"`
	Documentattributes interface{} `json:"DocumentAttributes,omitempty"`
	Documentexcerpt TextWithHighlights `json:"DocumentExcerpt,omitempty"` // Provides text and information about where to highlight the text.
	TypeField interface{} `json:"Type,omitempty"`
}

// UpdateThesaurusRequest represents the UpdateThesaurusRequest schema from the OpenAPI specification
type UpdateThesaurusRequest struct {
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Sources3path S3Path `json:"SourceS3Path,omitempty"` // Information required to find a specific file in an Amazon S3 bucket.
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id"`
}

// ListIndicesResponse represents the ListIndicesResponse schema from the OpenAPI specification
type ListIndicesResponse struct {
	Indexconfigurationsummaryitems interface{} `json:"IndexConfigurationSummaryItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeQuerySuggestionsBlockListRequest represents the DescribeQuerySuggestionsBlockListRequest schema from the OpenAPI specification
type DescribeQuerySuggestionsBlockListRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// ListAccessControlConfigurationsRequest represents the ListAccessControlConfigurationsRequest schema from the OpenAPI specification
type ListAccessControlConfigurationsRequest struct {
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AclConfiguration represents the AclConfiguration schema from the OpenAPI specification
type AclConfiguration struct {
	Allowedgroupscolumnname interface{} `json:"AllowedGroupsColumnName"`
}

// Correction represents the Correction schema from the OpenAPI specification
type Correction struct {
	Beginoffset interface{} `json:"BeginOffset,omitempty"`
	Correctedterm interface{} `json:"CorrectedTerm,omitempty"`
	Endoffset interface{} `json:"EndOffset,omitempty"`
	Term interface{} `json:"Term,omitempty"`
}

// CreateIndexResponse represents the CreateIndexResponse schema from the OpenAPI specification
type CreateIndexResponse struct {
	Id interface{} `json:"Id,omitempty"`
}

// DeleteQuerySuggestionsBlockListRequest represents the DeleteQuerySuggestionsBlockListRequest schema from the OpenAPI specification
type DeleteQuerySuggestionsBlockListRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// QueryRequest represents the QueryRequest schema from the OpenAPI specification
type QueryRequest struct {
	Indexid interface{} `json:"IndexId"`
	Pagesize interface{} `json:"PageSize,omitempty"`
	Querytext interface{} `json:"QueryText,omitempty"`
	Requesteddocumentattributes interface{} `json:"RequestedDocumentAttributes,omitempty"`
	Usercontext interface{} `json:"UserContext,omitempty"`
	Visitorid interface{} `json:"VisitorId,omitempty"`
	Queryresulttypefilter interface{} `json:"QueryResultTypeFilter,omitempty"`
	Sortingconfiguration interface{} `json:"SortingConfiguration,omitempty"`
	Spellcorrectionconfiguration interface{} `json:"SpellCorrectionConfiguration,omitempty"`
	Attributefilter interface{} `json:"AttributeFilter,omitempty"`
	Facets interface{} `json:"Facets,omitempty"`
	Pagenumber interface{} `json:"PageNumber,omitempty"`
	Documentrelevanceoverrideconfigurations interface{} `json:"DocumentRelevanceOverrideConfigurations,omitempty"`
}

// MemberUser represents the MemberUser schema from the OpenAPI specification
type MemberUser struct {
	Userid interface{} `json:"UserId"`
}

// EntityDisplayData represents the EntityDisplayData schema from the OpenAPI specification
type EntityDisplayData struct {
	Identifiedusername interface{} `json:"IdentifiedUserName,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// OneDriveConfiguration represents the OneDriveConfiguration schema from the OpenAPI specification
type OneDriveConfiguration struct {
	Secretarn interface{} `json:"SecretArn"`
	Tenantdomain interface{} `json:"TenantDomain"`
	Disablelocalgroups interface{} `json:"DisableLocalGroups,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Onedriveusers interface{} `json:"OneDriveUsers"`
}

// JwtTokenTypeConfiguration represents the JwtTokenTypeConfiguration schema from the OpenAPI specification
type JwtTokenTypeConfiguration struct {
	Url interface{} `json:"URL,omitempty"`
	Usernameattributefield interface{} `json:"UserNameAttributeField,omitempty"`
	Claimregex interface{} `json:"ClaimRegex,omitempty"`
	Groupattributefield interface{} `json:"GroupAttributeField,omitempty"`
	Issuer interface{} `json:"Issuer,omitempty"`
	Keylocation interface{} `json:"KeyLocation"`
	Secretmanagerarn interface{} `json:"SecretManagerArn,omitempty"`
}

// GetSnapshotsRequest represents the GetSnapshotsRequest schema from the OpenAPI specification
type GetSnapshotsRequest struct {
	Indexid interface{} `json:"IndexId"`
	Interval interface{} `json:"Interval"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Metrictype interface{} `json:"MetricType"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateFeaturedResultsSetResponse represents the UpdateFeaturedResultsSetResponse schema from the OpenAPI specification
type UpdateFeaturedResultsSetResponse struct {
	Featuredresultsset interface{} `json:"FeaturedResultsSet,omitempty"`
}

// SuggestionTextWithHighlights represents the SuggestionTextWithHighlights schema from the OpenAPI specification
type SuggestionTextWithHighlights struct {
	Highlights interface{} `json:"Highlights,omitempty"`
	Text interface{} `json:"Text,omitempty"`
}

// ListDataSourceSyncJobsResponse represents the ListDataSourceSyncJobsResponse schema from the OpenAPI specification
type ListDataSourceSyncJobsResponse struct {
	History interface{} `json:"History,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteExperienceResponse represents the DeleteExperienceResponse schema from the OpenAPI specification
type DeleteExperienceResponse struct {
}

// SiteMapsConfiguration represents the SiteMapsConfiguration schema from the OpenAPI specification
type SiteMapsConfiguration struct {
	Sitemaps interface{} `json:"SiteMaps"`
}

// StopDataSourceSyncJobRequest represents the StopDataSourceSyncJobRequest schema from the OpenAPI specification
type StopDataSourceSyncJobRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// DescribeFaqRequest represents the DescribeFaqRequest schema from the OpenAPI specification
type DescribeFaqRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// ListAccessControlConfigurationsResponse represents the ListAccessControlConfigurationsResponse schema from the OpenAPI specification
type ListAccessControlConfigurationsResponse struct {
	Accesscontrolconfigurations interface{} `json:"AccessControlConfigurations"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UserIdentityConfiguration represents the UserIdentityConfiguration schema from the OpenAPI specification
type UserIdentityConfiguration struct {
	Identityattributename interface{} `json:"IdentityAttributeName,omitempty"`
}

// CreateFaqResponse represents the CreateFaqResponse schema from the OpenAPI specification
type CreateFaqResponse struct {
	Id interface{} `json:"Id,omitempty"`
}

// SourceDocument represents the SourceDocument schema from the OpenAPI specification
type SourceDocument struct {
	Documentid interface{} `json:"DocumentId,omitempty"`
	Suggestionattributes interface{} `json:"SuggestionAttributes,omitempty"`
	Additionalattributes interface{} `json:"AdditionalAttributes,omitempty"`
}

// ListFaqsRequest represents the ListFaqsRequest schema from the OpenAPI specification
type ListFaqsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ListExperienceEntitiesResponse represents the ListExperienceEntitiesResponse schema from the OpenAPI specification
type ListExperienceEntitiesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Summaryitems interface{} `json:"SummaryItems,omitempty"`
}

// BasicAuthenticationConfiguration represents the BasicAuthenticationConfiguration schema from the OpenAPI specification
type BasicAuthenticationConfiguration struct {
	Credentials interface{} `json:"Credentials"`
	Host interface{} `json:"Host"`
	Port interface{} `json:"Port"`
}

// SpellCorrectionConfiguration represents the SpellCorrectionConfiguration schema from the OpenAPI specification
type SpellCorrectionConfiguration struct {
	Includequeryspellchecksuggestions interface{} `json:"IncludeQuerySpellCheckSuggestions"`
}

// ValueImportanceMap represents the ValueImportanceMap schema from the OpenAPI specification
type ValueImportanceMap struct {
}

// AssociatePersonasToEntitiesResponse represents the AssociatePersonasToEntitiesResponse schema from the OpenAPI specification
type AssociatePersonasToEntitiesResponse struct {
	Failedentitylist interface{} `json:"FailedEntityList,omitempty"`
}

// ListThesauriResponse represents the ListThesauriResponse schema from the OpenAPI specification
type ListThesauriResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Thesaurussummaryitems interface{} `json:"ThesaurusSummaryItems,omitempty"`
}

// BatchDeleteDocumentRequest represents the BatchDeleteDocumentRequest schema from the OpenAPI specification
type BatchDeleteDocumentRequest struct {
	Datasourcesyncjobmetrictarget DataSourceSyncJobMetricTarget `json:"DataSourceSyncJobMetricTarget,omitempty"` // Maps a particular data source sync job to a particular data source.
	Documentidlist interface{} `json:"DocumentIdList"`
	Indexid interface{} `json:"IndexId"`
}

// ConfluenceSpaceConfiguration represents the ConfluenceSpaceConfiguration schema from the OpenAPI specification
type ConfluenceSpaceConfiguration struct {
	Crawlarchivedspaces interface{} `json:"CrawlArchivedSpaces,omitempty"`
	Crawlpersonalspaces interface{} `json:"CrawlPersonalSpaces,omitempty"`
	Excludespaces interface{} `json:"ExcludeSpaces,omitempty"`
	Includespaces interface{} `json:"IncludeSpaces,omitempty"`
	Spacefieldmappings interface{} `json:"SpaceFieldMappings,omitempty"`
}

// DeleteAccessControlConfigurationResponse represents the DeleteAccessControlConfigurationResponse schema from the OpenAPI specification
type DeleteAccessControlConfigurationResponse struct {
}

// DeletePrincipalMappingRequest represents the DeletePrincipalMappingRequest schema from the OpenAPI specification
type DeletePrincipalMappingRequest struct {
	Orderingid interface{} `json:"OrderingId,omitempty"`
	Datasourceid interface{} `json:"DataSourceId,omitempty"`
	Groupid interface{} `json:"GroupId"`
	Indexid interface{} `json:"IndexId"`
}

// UpdateQuerySuggestionsConfigRequest represents the UpdateQuerySuggestionsConfigRequest schema from the OpenAPI specification
type UpdateQuerySuggestionsConfigRequest struct {
	Queryloglookbackwindowindays interface{} `json:"QueryLogLookBackWindowInDays,omitempty"`
	Attributesuggestionsconfig interface{} `json:"AttributeSuggestionsConfig,omitempty"`
	Includequerieswithoutuserinformation interface{} `json:"IncludeQueriesWithoutUserInformation,omitempty"`
	Indexid interface{} `json:"IndexId"`
	Minimumnumberofqueryingusers interface{} `json:"MinimumNumberOfQueryingUsers,omitempty"`
	Minimumquerycount interface{} `json:"MinimumQueryCount,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
}

// DocumentAttributeValue represents the DocumentAttributeValue schema from the OpenAPI specification
type DocumentAttributeValue struct {
	Longvalue interface{} `json:"LongValue,omitempty"`
	Stringlistvalue interface{} `json:"StringListValue,omitempty"`
	Stringvalue interface{} `json:"StringValue,omitempty"`
	Datevalue interface{} `json:"DateValue,omitempty"`
}

// CustomDocumentEnrichmentConfiguration represents the CustomDocumentEnrichmentConfiguration schema from the OpenAPI specification
type CustomDocumentEnrichmentConfiguration struct {
	Inlineconfigurations interface{} `json:"InlineConfigurations,omitempty"`
	Postextractionhookconfiguration interface{} `json:"PostExtractionHookConfiguration,omitempty"`
	Preextractionhookconfiguration interface{} `json:"PreExtractionHookConfiguration,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
}

// AttributeSuggestionsDescribeConfig represents the AttributeSuggestionsDescribeConfig schema from the OpenAPI specification
type AttributeSuggestionsDescribeConfig struct {
	Attributesuggestionsmode interface{} `json:"AttributeSuggestionsMode,omitempty"`
	Suggestableconfiglist interface{} `json:"SuggestableConfigList,omitempty"`
}

// ListExperiencesResponse represents the ListExperiencesResponse schema from the OpenAPI specification
type ListExperiencesResponse struct {
	Summaryitems interface{} `json:"SummaryItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AssociateEntitiesToExperienceResponse represents the AssociateEntitiesToExperienceResponse schema from the OpenAPI specification
type AssociateEntitiesToExperienceResponse struct {
	Failedentitylist interface{} `json:"FailedEntityList,omitempty"`
}

// HookConfiguration represents the HookConfiguration schema from the OpenAPI specification
type HookConfiguration struct {
	Invocationcondition interface{} `json:"InvocationCondition,omitempty"`
	Lambdaarn interface{} `json:"LambdaArn"`
	S3bucket interface{} `json:"S3Bucket"`
}

// Template represents the Template schema from the OpenAPI specification
type Template struct {
}

// DeleteExperienceRequest represents the DeleteExperienceRequest schema from the OpenAPI specification
type DeleteExperienceRequest struct {
	Indexid interface{} `json:"IndexId"`
	Id interface{} `json:"Id"`
}

// AlfrescoConfiguration represents the AlfrescoConfiguration schema from the OpenAPI specification
type AlfrescoConfiguration struct {
	Crawlcomments interface{} `json:"CrawlComments,omitempty"`
	Crawlsystemfolders interface{} `json:"CrawlSystemFolders,omitempty"`
	Entityfilter interface{} `json:"EntityFilter,omitempty"`
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Siteid interface{} `json:"SiteId"`
	Siteurl interface{} `json:"SiteUrl"`
	Sslcertificates3path interface{} `json:"SslCertificateS3Path"`
	Blogfieldmappings interface{} `json:"BlogFieldMappings,omitempty"`
	Documentlibraryfieldmappings interface{} `json:"DocumentLibraryFieldMappings,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Wikifieldmappings interface{} `json:"WikiFieldMappings,omitempty"`
}

// EntityConfiguration represents the EntityConfiguration schema from the OpenAPI specification
type EntityConfiguration struct {
	Entityid interface{} `json:"EntityId"`
	Entitytype interface{} `json:"EntityType"`
}

// TableRow represents the TableRow schema from the OpenAPI specification
type TableRow struct {
	Cells interface{} `json:"Cells,omitempty"`
}

// SaaSConfiguration represents the SaaSConfiguration schema from the OpenAPI specification
type SaaSConfiguration struct {
	Hosturl interface{} `json:"HostUrl"`
	Organizationname interface{} `json:"OrganizationName"`
}

// ConfluencePageConfiguration represents the ConfluencePageConfiguration schema from the OpenAPI specification
type ConfluencePageConfiguration struct {
	Pagefieldmappings interface{} `json:"PageFieldMappings,omitempty"`
}

// DocumentAttributeTarget represents the DocumentAttributeTarget schema from the OpenAPI specification
type DocumentAttributeTarget struct {
	Targetdocumentattributekey interface{} `json:"TargetDocumentAttributeKey,omitempty"`
	Targetdocumentattributevalue interface{} `json:"TargetDocumentAttributeValue,omitempty"`
	Targetdocumentattributevaluedeletion interface{} `json:"TargetDocumentAttributeValueDeletion,omitempty"`
}

// Warning represents the Warning schema from the OpenAPI specification
type Warning struct {
	Code interface{} `json:"Code,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// DescribeQuerySuggestionsConfigRequest represents the DescribeQuerySuggestionsConfigRequest schema from the OpenAPI specification
type DescribeQuerySuggestionsConfigRequest struct {
	Indexid interface{} `json:"IndexId"`
}

// EntityPersonaConfiguration represents the EntityPersonaConfiguration schema from the OpenAPI specification
type EntityPersonaConfiguration struct {
	Entityid interface{} `json:"EntityId"`
	Persona interface{} `json:"Persona"`
}

// Search represents the Search schema from the OpenAPI specification
type Search struct {
	Facetable interface{} `json:"Facetable,omitempty"`
	Searchable interface{} `json:"Searchable,omitempty"`
	Sortable interface{} `json:"Sortable,omitempty"`
	Displayable interface{} `json:"Displayable,omitempty"`
}

// CreateQuerySuggestionsBlockListRequest represents the CreateQuerySuggestionsBlockListRequest schema from the OpenAPI specification
type CreateQuerySuggestionsBlockListRequest struct {
	Indexid interface{} `json:"IndexId"`
	Name interface{} `json:"Name"`
	Rolearn interface{} `json:"RoleArn"`
	Sources3path interface{} `json:"SourceS3Path"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// SqlConfiguration represents the SqlConfiguration schema from the OpenAPI specification
type SqlConfiguration struct {
	Queryidentifiersenclosingoption interface{} `json:"QueryIdentifiersEnclosingOption,omitempty"`
}

// S3DataSourceConfiguration represents the S3DataSourceConfiguration schema from the OpenAPI specification
type S3DataSourceConfiguration struct {
	Documentsmetadataconfiguration DocumentsMetadataConfiguration `json:"DocumentsMetadataConfiguration,omitempty"` // Document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes. Each metadata file contains metadata about a single document.
	Exclusionpatterns interface{} `json:"ExclusionPatterns,omitempty"`
	Inclusionpatterns interface{} `json:"InclusionPatterns,omitempty"`
	Inclusionprefixes interface{} `json:"InclusionPrefixes,omitempty"`
	Accesscontrollistconfiguration interface{} `json:"AccessControlListConfiguration,omitempty"`
	Bucketname interface{} `json:"BucketName"`
}

// DocumentInfo represents the DocumentInfo schema from the OpenAPI specification
type DocumentInfo struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Documentid interface{} `json:"DocumentId"`
}

// ListDataSourcesResponse represents the ListDataSourcesResponse schema from the OpenAPI specification
type ListDataSourcesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Summaryitems interface{} `json:"SummaryItems,omitempty"`
}

// DataSourceGroup represents the DataSourceGroup schema from the OpenAPI specification
type DataSourceGroup struct {
	Groupid interface{} `json:"GroupId"`
	Datasourceid interface{} `json:"DataSourceId"`
}

// SortingConfiguration represents the SortingConfiguration schema from the OpenAPI specification
type SortingConfiguration struct {
	Sortorder interface{} `json:"SortOrder"`
	Documentattributekey interface{} `json:"DocumentAttributeKey"`
}

// BatchPutDocumentRequest represents the BatchPutDocumentRequest schema from the OpenAPI specification
type BatchPutDocumentRequest struct {
	Customdocumentenrichmentconfiguration interface{} `json:"CustomDocumentEnrichmentConfiguration,omitempty"`
	Documents interface{} `json:"Documents"`
	Indexid interface{} `json:"IndexId"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
}

// AdditionalResultAttributeValue represents the AdditionalResultAttributeValue schema from the OpenAPI specification
type AdditionalResultAttributeValue struct {
	Textwithhighlightsvalue interface{} `json:"TextWithHighlightsValue,omitempty"`
}

// DocumentAttributeValueCountPair represents the DocumentAttributeValueCountPair schema from the OpenAPI specification
type DocumentAttributeValueCountPair struct {
	Documentattributevalue interface{} `json:"DocumentAttributeValue,omitempty"`
	Facetresults interface{} `json:"FacetResults,omitempty"`
	Count interface{} `json:"Count,omitempty"`
}

// BatchPutDocumentResponse represents the BatchPutDocumentResponse schema from the OpenAPI specification
type BatchPutDocumentResponse struct {
	Faileddocuments interface{} `json:"FailedDocuments,omitempty"`
}

// AttributeSuggestionsGetConfig represents the AttributeSuggestionsGetConfig schema from the OpenAPI specification
type AttributeSuggestionsGetConfig struct {
	Additionalresponseattributes interface{} `json:"AdditionalResponseAttributes,omitempty"`
	Attributefilter interface{} `json:"AttributeFilter,omitempty"`
	Suggestionattributes interface{} `json:"SuggestionAttributes,omitempty"`
	Usercontext interface{} `json:"UserContext,omitempty"`
}

// DataSourceSyncJob represents the DataSourceSyncJob schema from the OpenAPI specification
type DataSourceSyncJob struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
	Metrics interface{} `json:"Metrics,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Datasourceerrorcode interface{} `json:"DataSourceErrorCode,omitempty"`
}

// SubmitFeedbackRequest represents the SubmitFeedbackRequest schema from the OpenAPI specification
type SubmitFeedbackRequest struct {
	Indexid interface{} `json:"IndexId"`
	Queryid interface{} `json:"QueryId"`
	Relevancefeedbackitems interface{} `json:"RelevanceFeedbackItems,omitempty"`
	Clickfeedbackitems interface{} `json:"ClickFeedbackItems,omitempty"`
}

// DocumentAttribute represents the DocumentAttribute schema from the OpenAPI specification
type DocumentAttribute struct {
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// ExperienceConfiguration represents the ExperienceConfiguration schema from the OpenAPI specification
type ExperienceConfiguration struct {
	Useridentityconfiguration interface{} `json:"UserIdentityConfiguration,omitempty"`
	Contentsourceconfiguration interface{} `json:"ContentSourceConfiguration,omitempty"`
}

// DataSourceSyncJobMetrics represents the DataSourceSyncJobMetrics schema from the OpenAPI specification
type DataSourceSyncJobMetrics struct {
	Documentsmodified interface{} `json:"DocumentsModified,omitempty"`
	Documentsscanned interface{} `json:"DocumentsScanned,omitempty"`
	Documentsadded interface{} `json:"DocumentsAdded,omitempty"`
	Documentsdeleted interface{} `json:"DocumentsDeleted,omitempty"`
	Documentsfailed interface{} `json:"DocumentsFailed,omitempty"`
}

// CreateAccessControlConfigurationResponse represents the CreateAccessControlConfigurationResponse schema from the OpenAPI specification
type CreateAccessControlConfigurationResponse struct {
	Id interface{} `json:"Id"`
}

// UpdateAccessControlConfigurationResponse represents the UpdateAccessControlConfigurationResponse schema from the OpenAPI specification
type UpdateAccessControlConfigurationResponse struct {
}

// PersonasSummary represents the PersonasSummary schema from the OpenAPI specification
type PersonasSummary struct {
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Entityid interface{} `json:"EntityId,omitempty"`
	Persona interface{} `json:"Persona,omitempty"`
}

// TimeRange represents the TimeRange schema from the OpenAPI specification
type TimeRange struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tags interface{} `json:"Tags"`
}

// BatchGetDocumentStatusResponseError represents the BatchGetDocumentStatusResponseError schema from the OpenAPI specification
type BatchGetDocumentStatusResponseError struct {
	Documentid interface{} `json:"DocumentId,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// Relevance represents the Relevance schema from the OpenAPI specification
type Relevance struct {
	Duration interface{} `json:"Duration,omitempty"`
	Freshness interface{} `json:"Freshness,omitempty"`
	Importance interface{} `json:"Importance,omitempty"`
	Rankorder interface{} `json:"RankOrder,omitempty"`
	Valueimportancemap interface{} `json:"ValueImportanceMap,omitempty"`
}

// ServiceNowConfiguration represents the ServiceNowConfiguration schema from the OpenAPI specification
type ServiceNowConfiguration struct {
	Authenticationtype interface{} `json:"AuthenticationType,omitempty"`
	Hosturl interface{} `json:"HostUrl"`
	Knowledgearticleconfiguration interface{} `json:"KnowledgeArticleConfiguration,omitempty"`
	Secretarn interface{} `json:"SecretArn"`
	Servicecatalogconfiguration interface{} `json:"ServiceCatalogConfiguration,omitempty"`
	Servicenowbuildversion interface{} `json:"ServiceNowBuildVersion"`
}

// DescribeThesaurusRequest represents the DescribeThesaurusRequest schema from the OpenAPI specification
type DescribeThesaurusRequest struct {
	Id interface{} `json:"Id"`
	Indexid interface{} `json:"IndexId"`
}

// SalesforceCustomKnowledgeArticleTypeConfiguration represents the SalesforceCustomKnowledgeArticleTypeConfiguration schema from the OpenAPI specification
type SalesforceCustomKnowledgeArticleTypeConfiguration struct {
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Name interface{} `json:"Name"`
	Documentdatafieldname interface{} `json:"DocumentDataFieldName"`
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
}

// DatabaseConfiguration represents the DatabaseConfiguration schema from the OpenAPI specification
type DatabaseConfiguration struct {
	Columnconfiguration interface{} `json:"ColumnConfiguration"`
	Connectionconfiguration interface{} `json:"ConnectionConfiguration"`
	Databaseenginetype interface{} `json:"DatabaseEngineType"`
	Sqlconfiguration interface{} `json:"SqlConfiguration,omitempty"`
	Vpcconfiguration DataSourceVpcConfiguration `json:"VpcConfiguration,omitempty"` // Provides the configuration information to connect to an Amazon VPC.
	Aclconfiguration interface{} `json:"AclConfiguration,omitempty"`
}

// BatchDeleteFeaturedResultsSetError represents the BatchDeleteFeaturedResultsSetError schema from the OpenAPI specification
type BatchDeleteFeaturedResultsSetError struct {
	Errorcode interface{} `json:"ErrorCode"`
	Errormessage interface{} `json:"ErrorMessage"`
	Id interface{} `json:"Id"`
}

// CreateQuerySuggestionsBlockListResponse represents the CreateQuerySuggestionsBlockListResponse schema from the OpenAPI specification
type CreateQuerySuggestionsBlockListResponse struct {
	Id interface{} `json:"Id,omitempty"`
}

// DescribePrincipalMappingResponse represents the DescribePrincipalMappingResponse schema from the OpenAPI specification
type DescribePrincipalMappingResponse struct {
	Grouporderingidsummaries interface{} `json:"GroupOrderingIdSummaries,omitempty"`
	Indexid interface{} `json:"IndexId,omitempty"`
	Datasourceid interface{} `json:"DataSourceId,omitempty"`
	Groupid interface{} `json:"GroupId,omitempty"`
}

// BatchDeleteDocumentResponseFailedDocument represents the BatchDeleteDocumentResponseFailedDocument schema from the OpenAPI specification
type BatchDeleteDocumentResponseFailedDocument struct {
	Id interface{} `json:"Id,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// BatchGetDocumentStatusResponse represents the BatchGetDocumentStatusResponse schema from the OpenAPI specification
type BatchGetDocumentStatusResponse struct {
	Documentstatuslist interface{} `json:"DocumentStatusList,omitempty"`
	Errors interface{} `json:"Errors,omitempty"`
}

// TextDocumentStatistics represents the TextDocumentStatistics schema from the OpenAPI specification
type TextDocumentStatistics struct {
	Indexedtextdocumentscount interface{} `json:"IndexedTextDocumentsCount"`
	Indexedtextbytes interface{} `json:"IndexedTextBytes"`
}

// ListIndicesRequest represents the ListIndicesRequest schema from the OpenAPI specification
type ListIndicesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateIndexRequest represents the UpdateIndexRequest schema from the OpenAPI specification
type UpdateIndexRequest struct {
	Usertokenconfigurations interface{} `json:"UserTokenConfigurations,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Capacityunits interface{} `json:"CapacityUnits,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Documentmetadataconfigurationupdates interface{} `json:"DocumentMetadataConfigurationUpdates,omitempty"`
	Id interface{} `json:"Id"`
	Name interface{} `json:"Name,omitempty"`
	Usercontextpolicy interface{} `json:"UserContextPolicy,omitempty"`
	Usergroupresolutionconfiguration interface{} `json:"UserGroupResolutionConfiguration,omitempty"`
}

// RetrieveResult represents the RetrieveResult schema from the OpenAPI specification
type RetrieveResult struct {
	Queryid interface{} `json:"QueryId,omitempty"`
	Resultitems interface{} `json:"ResultItems,omitempty"`
}

// DescribeExperienceResponse represents the DescribeExperienceResponse schema from the OpenAPI specification
type DescribeExperienceResponse struct {
	Description interface{} `json:"Description,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Indexid interface{} `json:"IndexId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// StartDataSourceSyncJobResponse represents the StartDataSourceSyncJobResponse schema from the OpenAPI specification
type StartDataSourceSyncJobResponse struct {
	Executionid interface{} `json:"ExecutionId,omitempty"`
}

// SalesforceKnowledgeArticleConfiguration represents the SalesforceKnowledgeArticleConfiguration schema from the OpenAPI specification
type SalesforceKnowledgeArticleConfiguration struct {
	Customknowledgearticletypeconfigurations interface{} `json:"CustomKnowledgeArticleTypeConfigurations,omitempty"`
	Includedstates interface{} `json:"IncludedStates"`
	Standardknowledgearticletypeconfiguration interface{} `json:"StandardKnowledgeArticleTypeConfiguration,omitempty"`
}

// HierarchicalPrincipal represents the HierarchicalPrincipal schema from the OpenAPI specification
type HierarchicalPrincipal struct {
	Principallist interface{} `json:"PrincipalList"`
}

// AttributeSuggestionsUpdateConfig represents the AttributeSuggestionsUpdateConfig schema from the OpenAPI specification
type AttributeSuggestionsUpdateConfig struct {
	Attributesuggestionsmode interface{} `json:"AttributeSuggestionsMode,omitempty"`
	Suggestableconfiglist interface{} `json:"SuggestableConfigList,omitempty"`
}

// DocumentRelevanceConfiguration represents the DocumentRelevanceConfiguration schema from the OpenAPI specification
type DocumentRelevanceConfiguration struct {
	Name interface{} `json:"Name"`
	Relevance interface{} `json:"Relevance"`
}

// DataSourceSummary represents the DataSourceSummary schema from the OpenAPI specification
type DataSourceSummary struct {
	Id interface{} `json:"Id,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// FaqStatistics represents the FaqStatistics schema from the OpenAPI specification
type FaqStatistics struct {
	Indexedquestionanswerscount interface{} `json:"IndexedQuestionAnswersCount"`
}

// SalesforceChatterFeedConfiguration represents the SalesforceChatterFeedConfiguration schema from the OpenAPI specification
type SalesforceChatterFeedConfiguration struct {
	Documentdatafieldname interface{} `json:"DocumentDataFieldName"`
	Documenttitlefieldname interface{} `json:"DocumentTitleFieldName,omitempty"`
	Fieldmappings interface{} `json:"FieldMappings,omitempty"`
	Includefiltertypes interface{} `json:"IncludeFilterTypes,omitempty"`
}

// FacetResult represents the FacetResult schema from the OpenAPI specification
type FacetResult struct {
	Documentattributekey interface{} `json:"DocumentAttributeKey,omitempty"`
	Documentattributevaluecountpairs interface{} `json:"DocumentAttributeValueCountPairs,omitempty"`
	Documentattributevaluetype interface{} `json:"DocumentAttributeValueType,omitempty"`
}

// RetrieveRequest represents the RetrieveRequest schema from the OpenAPI specification
type RetrieveRequest struct {
	Pagenumber interface{} `json:"PageNumber,omitempty"`
	Pagesize interface{} `json:"PageSize,omitempty"`
	Querytext interface{} `json:"QueryText"`
	Requesteddocumentattributes interface{} `json:"RequestedDocumentAttributes,omitempty"`
	Usercontext interface{} `json:"UserContext,omitempty"`
	Attributefilter interface{} `json:"AttributeFilter,omitempty"`
	Documentrelevanceoverrideconfigurations interface{} `json:"DocumentRelevanceOverrideConfigurations,omitempty"`
	Indexid interface{} `json:"IndexId"`
}

// IndexStatistics represents the IndexStatistics schema from the OpenAPI specification
type IndexStatistics struct {
	Faqstatistics interface{} `json:"FaqStatistics"`
	Textdocumentstatistics interface{} `json:"TextDocumentStatistics"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Failurecode interface{} `json:"FailureCode,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Documentid interface{} `json:"DocumentId,omitempty"`
	Documentstatus interface{} `json:"DocumentStatus,omitempty"`
}

// OnPremiseConfiguration represents the OnPremiseConfiguration schema from the OpenAPI specification
type OnPremiseConfiguration struct {
	Sslcertificates3path interface{} `json:"SslCertificateS3Path"`
	Hosturl interface{} `json:"HostUrl"`
	Organizationname interface{} `json:"OrganizationName"`
}

// CreateDataSourceResponse represents the CreateDataSourceResponse schema from the OpenAPI specification
type CreateDataSourceResponse struct {
	Id interface{} `json:"Id"`
}

// AccessControlListConfiguration represents the AccessControlListConfiguration schema from the OpenAPI specification
type AccessControlListConfiguration struct {
	Keypath interface{} `json:"KeyPath,omitempty"`
}
