<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="股票列表核心表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                股票代码 (唯一业务标识，如: 000001)
              </template>
              {{ formValue.dm }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                股票名称 (如: 平安银行)
              </template>
              {{ formValue.mc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易所代码 (如: sh, sz, bj)
              </template>
              {{ formValue.jys }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易所全称
              </template>
              {{ formValue.exchangeName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标准股票代码 (如: 000001.SZ)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item label="状态: 1-正常, 0-退市">
              <n-tag
            :type="dict.getType('sys_normal_disable', formValue.status)"
            size="small"
            class="min-left-space"
          >
                {{ dict.getLabel('sys_normal_disable', formValue.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                上市日期
              </template>
              {{ formValue.listDate }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/stockList';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useDictStore } from '@/store/modules/dict';

  const message = useMessage();
  const dict = useDictStore();
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