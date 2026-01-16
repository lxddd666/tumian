<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="资金流向明细表详情" closable>
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
                交易时间 (通常为HHMMSS格式的整数)
              </template>
              {{ formValue.t }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买单总单数
              </template>
              {{ formValue.zmbzds }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖单总单数
              </template>
              {{ formValue.zmszds }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                大单动向
              </template>
              {{ formValue.dddx }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                涨跌动因
              </template>
              {{ formValue.zddy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                大单差分
              </template>
              {{ formValue.ddcf }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买单总单数增量
              </template>
              {{ formValue.zmbzdszl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖单总单数增量
              </template>
              {{ formValue.zmszdszl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                成交笔数增量
              </template>
              {{ formValue.cjbszl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买特大单成交额
              </template>
              {{ formValue.zmbtdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买大单成交额
              </template>
              {{ formValue.zmbddcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买中单成交额
              </template>
              {{ formValue.zmbzdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买小单成交额
              </template>
              {{ formValue.zmbxdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖特大单成交额
              </template>
              {{ formValue.zmstdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖大单成交额
              </template>
              {{ formValue.zmsddcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖中单成交额
              </template>
              {{ formValue.zmszdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖小单成交额
              </template>
              {{ formValue.zmsxdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买特大单成交额
              </template>
              {{ formValue.bdmbtdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买大单成交额
              </template>
              {{ formValue.bdmbddcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买中单成交额
              </template>
              {{ formValue.bdmbzdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买小单成交额
              </template>
              {{ formValue.bdmbxdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖特大单成交额
              </template>
              {{ formValue.bdmstdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖大单成交额
              </template>
              {{ formValue.bdmsddcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖中单成交额
              </template>
              {{ formValue.bdmszdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖小单成交额
              </template>
              {{ formValue.bdmsxdcje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买特大单成交量
              </template>
              {{ formValue.zmbtdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买大单成交量
              </template>
              {{ formValue.zmbddcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买中单成交量
              </template>
              {{ formValue.zmbzdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买小单成交量
              </template>
              {{ formValue.zmbxdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖特大单成交量
              </template>
              {{ formValue.zmstdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖大单成交量
              </template>
              {{ formValue.zmsddcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖中单成交量
              </template>
              {{ formValue.zmszdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主卖小单成交量
              </template>
              {{ formValue.zmsxdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买特大单成交量
              </template>
              {{ formValue.bdmbtdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买大单成交量
              </template>
              {{ formValue.bdmbddcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买中单成交量
              </template>
              {{ formValue.bdmbzdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动买小单成交量
              </template>
              {{ formValue.bdmbxdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖特大单成交量
              </template>
              {{ formValue.bdmstdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖大单成交量
              </template>
              {{ formValue.bdmsddcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖中单成交量
              </template>
              {{ formValue.bdmszdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被动卖小单成交量
              </template>
              {{ formValue.bdmsxdcjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买特大单成交额增量
              </template>
              {{ formValue.zmbtdcjzl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买大单成交额增量
              </template>
              {{ formValue.zmbddcjzl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买特大单成交量增量
              </template>
              {{ formValue.zmbtdcjzlv }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主买大单成交量增量
              </template>
              {{ formValue.zmbddcjzlv }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/flowOfFunds';
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