<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑企业级历史行情数据表 (K线数据) #' + formValue.id : '添加企业级历史行情数据表 (K线数据)'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid
              cols="1 s:1 m:1 l:1 xl:1 2xl:1"
              responsive="screen"
            >
              <n-gi span="1">
                <n-form-item label="证券代码 (如: 000001.SZ, AAPL)" path="symbol">
                  <n-input
                    placeholder="请输入证券代码 (如: 000001.SZ, AAPL)"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易时间 (精确到分钟或日)" path="t">
                  <DatePicker v-model:formValue="formValue.t" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易日期 (衍生字段)" path="date">
                  <DatePicker v-model:formValue="formValue.date" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易年份" path="year">
                  <n-input-number
                    placeholder="请输入交易年份"
                    v-model:value="formValue.year"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易月份" path="month">
                  <n-input-number
                    placeholder="请输入交易月份"
                    v-model:value="formValue.month"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="星期几 (1=周日,7=周六)" path="weekday">
                  <n-input-number
                    placeholder="请输入星期几 (1=周日,7=周六)"
                    v-model:value="formValue.weekday"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="开盘价" path="o">
                  <n-input-number
                    placeholder="请输入开盘价"
                    v-model:value="formValue.o"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最高价" path="h">
                  <n-input-number
                    placeholder="请输入最高价"
                    v-model:value="formValue.h"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最低价" path="l">
                  <n-input-number
                    placeholder="请输入最低价"
                    v-model:value="formValue.l"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="收盘价" path="c">
                  <n-input-number
                    placeholder="请输入收盘价"
                    v-model:value="formValue.c"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="前收盘价" path="pc">
                  <n-input-number
                    placeholder="请输入前收盘价"
                    v-model:value="formValue.pc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="成交量 (股/手)" path="v">
                  <n-input-number
                    placeholder="请输入成交量 (股/手)"
                    v-model:value="formValue.v"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="成交额 (元)" path="a">
                  <n-input-number
                    placeholder="请输入成交额 (元)"
                    v-model:value="formValue.a"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="涨跌额" path="change">
                  <n-input-number
                    placeholder="请输入涨跌额"
                    v-model:value="formValue.change"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="涨跌幅 (%)" path="changePct">
                  <n-input-number
                    placeholder="请输入涨跌幅 (%)"
                    v-model:value="formValue.changePct"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="振幅 (%)" path="amplitude">
                  <n-input-number
                    placeholder="请输入振幅 (%)"
                    v-model:value="formValue.amplitude"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="停牌标志: 0-正常, 1-停牌" path="sf">
                  <n-input-number
                    placeholder="请输入停牌标志: 0-正常, 1-停牌"
                    v-model:value="formValue.sf"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易状态描述" path="tradingStatus">
                  <n-input
                    placeholder="请输入交易状态描述"
                    v-model:value="formValue.tradingStatus"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month" path="period">
                  <n-input
                    placeholder="请输入数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month"
                    v-model:value="formValue.period"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否复权: 0-不复权, 1-前复权, 2-后复权" path="isAdjusted">
                  <n-input-number
                    placeholder="请输入是否复权: 0-不复权, 1-前复权, 2-后复权"
                    v-model:value="formValue.isAdjusted"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据质量: 0-异常, 1-正常, 2-补全" path="dataQuality">
                  <n-input-number
                    placeholder="请输入数据质量: 0-异常, 1-正常, 2-补全"
                    v-model:value="formValue.dataQuality"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否已验证: 0-未验证, 1-已验证" path="isVerified">
                  <n-input-number
                    placeholder="请输入是否已验证: 0-未验证, 1-已验证"
                    v-model:value="formValue.isVerified"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据来源" path="dataSource">
                  <n-input
                    placeholder="请输入数据来源"
                    v-model:value="formValue.dataSource"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据版本" path="version">
                  <n-input-number
                    placeholder="请输入数据版本"
                    v-model:value="formValue.version"
                    />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" :disabled="!isFormValid" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>
<script lang="ts" setup>
  import { Edit, View } from '@/api/enterpriseHistoricalData';
  import { State, newState, rules } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });
  const isFormValid = ref(true);

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    
    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      
      return;
    }

    // 编辑
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

<style lang="less"></style>