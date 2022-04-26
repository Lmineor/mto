<template>
    <div class="formatjson">
        <nya-container title="json格式化">
            <vue-json-editor
                class="vue-json-editor"
                v-model="resultInfo"
                :showBtns="false"
                :mode="'code'"
                lang="zh"
                @json-change="onJsonChange"
                @json-save="onJsonSave"
                @has-error="onError"
            />
        </nya-container>
    </div>
</template>

<script>
import envs from '../env'
// import { vueJsonEditor } from 'vue-json-editor'
export default {
    // components: { vueJsonEditor },
    name: '格式化Json',
    head() {
        return{
            title: 'json格式化'
        }
    },
    data() {
        return {
        hasJsonFlag:true,  // json是否验证通过
        // json数据
        resultInfo: {
          'employees': [
            {
            'firstName': 'Bill',
            'lastName': 'Gates'
            },
          ]
        }
      }
    },
    methods: {
        onJsonChange (value) {
        console.log('更改value:', value);
        // 实时保存
        this.onJsonSave(value)
      },
      onJsonSave (value) {
        // console.log('保存value:', value);
        this.resultInfo = value
        this.hasJsonFlag = true
      },
      onError(value) {
        console.log("json错误了value:", value);
        this.hasJsonFlag = false
      },
      // 检查json
      checkJson(){
        if (this.hasJsonFlag == false){
          // console.log("json验证失败")
          // this.$message.error("json验证失败")
          alert("json验证失败")
          return false
        } else {
          // console.log("json验证成功")
          // this.$message.success("json验证成功")
          alert("json验证成功")
          return true
        }
      },
    }
};
</script>

<style lang="scss">
.formatjson{
    height: 100vh !important;
    .vue-json-editor {
        height: 75vh !important;
    }
    .jsoneditor-vue {
        height: 100% !important;
    }
}
</style>
