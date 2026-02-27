import { request } from '../request';

/** get site list */
export function GetSiteList(params?: Api.Site.SiteSearchParams) {
  return request<Api.Site.SiteList>({
    url: '/api/v1/manage/getSiteList',
    method: 'get',
    params
  });
}

/** get all sites (enabled only) */
export function GetAllSites() {
  return request<Api.Site.AllSite[]>({
    url: '/api/v1/manage/getAllSites',
    method: 'get'
  });
}

/** add site */
export function AddSite(req: Api.Site.AddSiteRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/addSite',
    method: 'post',
    data: req
  });
}

/** edit site */
export function EditSite(req: Api.Site.EditSiteRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/editSite',
    method: 'post',
    data: req
  });
}

/** delete site */
export function DeleteSite(req: string[]) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/deleteSite',
    method: 'post',
    data: {
      uuids: req
    }
  });
}

/** get site topup list */
export function GetSiteTopupList(params?: Api.Site.SiteTopupSearchParams) {
  return request<Api.Site.SiteTopupList>({
    url: '/api/v1/manage/getSiteTopupList',
    method: 'get',
    params
  });
}

/** add site topup */
export function AddSiteTopup(req: Api.Site.AddSiteTopupRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/addSiteTopup',
    method: 'post',
    data: req
  });
}

/** get site deduction list */
export function GetSiteDeductionList(params?: Api.Site.SiteDeductionSearchParams) {
  return request<Api.Site.SiteDeductionList>({
    url: '/api/v1/manage/getSiteDeductionList',
    method: 'get',
    params
  });
}

/** add site deduction */
export function AddSiteDeduction(req: Api.Site.AddSiteDeductionRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/addSiteDeduction',
    method: 'post',
    data: req
  });
}

/** quick topup site score */
export function TopupSite(req: Api.Site.TopupSiteRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/topupSite',
    method: 'post',
    data: req
  });
}

/** quick deduction site score */
export function DeductionSite(req: Api.Site.DeductionSiteRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/deductionSite',
    method: 'post',
    data: req
  });
}

/** sync remote site remaining_score */
export function SyncRemoteSiteScore(req: Api.Site.SyncRemoteScoreRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/syncRemoteSiteScore',
    method: 'post',
    data: req
  });
}
