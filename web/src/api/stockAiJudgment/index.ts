import { http, jumpExport } from '@/utils/http/axios';

// 获取ai 选股判断列表
export function List(params) {
  return http.request({
    url: '/stockAiJudgment/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除ai 选股判断
export function Delete(params) {
  return http.request({
    url: '/stockAiJudgment/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑ai 选股判断
export function Edit(params) {
  return http.request({
    url: '/stockAiJudgment/edit',
    method: 'POST',
    params,
  });
}

// 获取ai 选股判断指定详情
export function View(params) {
  return http.request({
    url: '/stockAiJudgment/view',
    method: 'GET',
    params,
  });
}

// 导出ai 选股判断
export function Export(params) {
  jumpExport('/stockAiJudgment/export', params);
}