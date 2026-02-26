declare namespace Api {
  namespace Site {
    type CommonSearchParams = Pick<Common.PaginatingCommonParams, 'current' | 'size'>;

    type Site = Common.CommonRecord<{
      siteName: string;
      siteLogo: string;
      priceCurrency: string;
      scoreStatType: string;
      remark: string;
      dbHost: string;
      dbUsername: string;
      dbPassword: string;
      dbPort: number;
      dbName: string;
      jwtSecret: string;
      adminUrl: string;
      adminUsername: string;
      contactName: string;
      contactPhone: string;
      contactEmail: string;
      remainingScore: string;
      totalTopup: string;
    }>;

    type AddSiteRequest = {
      siteName: string;
      siteLogo: string;
      priceCurrency: string;
      scoreStatType: string;
      status: string | null;
      remark: string;
      dbHost: string;
      dbUsername: string;
      dbPassword: string;
      dbPort: number;
      dbName: string;
      jwtSecret: string;
      adminUrl: string;
      adminUsername: string;
      contactName: string;
      contactPhone: string;
      contactEmail: string;
    };

    type EditSiteRequest = {
      uuid: string | undefined;
      siteName: string;
      siteLogo: string;
      priceCurrency: string;
      scoreStatType: string;
      status: string | null;
      remark: string;
      dbHost: string;
      dbUsername: string;
      dbPassword: string;
      dbPort: number;
      dbName: string;
      jwtSecret: string;
      adminUrl: string;
      adminUsername: string;
      contactName: string;
      contactPhone: string;
      contactEmail: string;
    };

    type SiteSearchParams = CommonType.RecordNullable<
      Pick<Api.Site.Site, 'siteName' | 'priceCurrency' | 'status'> & CommonSearchParams
    >;

    type SiteList = Common.PaginatingQueryRecord<Site>;

    type AllSite = {
      uuid: string;
      siteName: string;
      priceCurrency: string;
    };

    type SiteTopup = {
      uuid: string;
      siteUuid: string;
      score: string;
      topupMethod: string;
      topupStatus: string;
      remark: string;
      siteName: string;
      priceCurrency: string;
      operatorUuid: string;
      createTime: string;
      updateTime: string;
    };

    type AddSiteTopupRequest = {
      siteUuid: string;
      score: string;
      topupMethod: string;
      topupStatus: string;
      remark: string;
    };

    type SiteTopupSearchParams = CommonType.RecordNullable<
      Pick<SiteTopup, 'siteUuid' | 'topupStatus' | 'siteName'> & CommonSearchParams
    >;

    type SiteTopupList = Common.PaginatingQueryRecord<SiteTopup>;

    type SiteDeduction = {
      uuid: string;
      siteUuid: string;
      score: string;
      deductionMethod: string;
      deductionStatus: string;
      remark: string;
      siteName: string;
      priceCurrency: string;
      operatorUuid: string;
      createTime: string;
      updateTime: string;
    };

    type AddSiteDeductionRequest = {
      siteUuid: string;
      score: string;
      deductionMethod: string;
      deductionStatus: string;
      remark: string;
    };

    type SiteDeductionSearchParams = CommonType.RecordNullable<
      Pick<SiteDeduction, 'siteUuid' | 'deductionStatus' | 'siteName'> & CommonSearchParams
    >;

    type SiteDeductionList = Common.PaginatingQueryRecord<SiteDeduction>;
  }
}
