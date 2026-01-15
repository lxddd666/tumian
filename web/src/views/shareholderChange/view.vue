<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
<n-drawer-content title="股东户数变化记录表 (记录相邻报告期的户数变化)详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                股票代码 (如: 000001.SZ)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                截止日期 (统计截止日，如2025-12-31)[citation:9]
              </template>
              {{ formValue.jzrq }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                股东户数 (统计截止日的总户数)[citation:3][citation:6]
              </template>
              {{ formValue.gdhs }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]
              </template>
              {{ formValue.bh }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                变化方向 (衍生字段)
              </template>
              {{ formValue.changeDirection }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                公告日期 (信息发布日期)[citation:1]
              </template>
              {{ formValue.annDate }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/shareholderChange';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';

  const message = useMessage();
  
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped></style>