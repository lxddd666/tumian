<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑公司十大流通股东表 (数据来源于定期报告)[citation:4] #' + formValue.id : '添加公司十大流通股东表 (数据来源于定期报告)[citation:4]'"
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
                <n-form-item label="公司代码/股票代码 (例如: 000001.SZ)" path="symbol">
                  <n-input
                    placeholder="请输入公司代码/股票代码 (例如: 000001.SZ)"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="截止日期 (报告期结束日, 如2023-09-30)[citation:4]" path="jzrq">
                  <DatePicker v-model:formValue="formValue.jzrq" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="公告日期 (信息发布日期)" path="ggrq">
                  <DatePicker v-model:formValue="formValue.ggrq" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告年度" path="reportYear">
                  <n-input-number
                    placeholder="请输入报告年度"
                    v-model:value="formValue.reportYear"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告季度 (1-4)" path="reportQuarter">
                  <n-input-number
                    placeholder="请输入报告季度 (1-4)"
                    v-model:value="formValue.reportQuarter"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
<n-form-item label="报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]" path="reportType">
                  <n-input
                    placeholder="请输入报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]"
                    v-model:value="formValue.reportType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="股东名称" path="gdmc">
                  <n-input
                    placeholder="请输入股东名称"
                    v-model:value="formValue.gdmc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="股东类型 (如: 基金、社保、个人等)" path="gdlx">
                  <n-input
                    placeholder="请输入股东类型 (如: 基金、社保、个人等)"
                    v-model:value="formValue.gdlx"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="股份性质 (如: 流通A股、限售A股等)" path="gfxz">
                  <n-input
                    placeholder="请输入股份性质 (如: 流通A股、限售A股等)"
                    v-model:value="formValue.gfxz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="持股数量 (股)" path="cgsl">
                  <n-input-number
                    placeholder="请输入持股数量 (股)"
                    v-model:value="formValue.cgsl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="持股比例 (%)" path="cgbl">
                  <n-input-number
                    placeholder="请输入持股比例 (%)"
                    v-model:value="formValue.cgbl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="持股排名 (1-10)" path="cgpm">
                  <n-input-number
                    placeholder="请输入持股排名 (1-10)"
                    v-model:value="formValue.cgpm"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="变动原因" path="bdyy">
                  <n-input
                    type="textarea"
                    placeholder="变动原因"
                    v-model:value="formValue.bdyy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
<n-form-item label="变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他" path="bdType">
                  <n-input
                    placeholder="请输入变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他"
                    v-model:value="formValue.bdType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据来源 (如: 交易所公告)[citation:4]" path="dataSource">
                  <n-input
                    placeholder="请输入数据来源 (如: 交易所公告)[citation:4]"
                    v-model:value="formValue.dataSource"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否为该报告期最新数据: 0-历史快照, 1-最新" path="isLatest">
                  <n-input-number
                    placeholder="请输入是否为该报告期最新数据: 0-历史快照, 1-最新"
                    v-model:value="formValue.isLatest"
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
  import { Edit, View } from '@/api/topTenCirculatingHolders';
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