<template>
  <div class="container">
    <el-button type="primary" @click="dialogVisible = true">新建公告</el-button>
    <el-table highlight-current-row="true" :data="tableData" style="width: 100%;margin-top:30px;" border>
      <el-table-column prop="date" label="发布日期" width="100">
      </el-table-column>
      <el-table-column prop="name" label="撰写人" width="100">
      </el-table-column>
      <el-table-column prop="bumen" label="所属部门" width="150">
      </el-table-column>
      <el-table-column prop="title" label="标题" width="200">
      </el-table-column>
      <el-table-column prop="importance" label="优先级" width="100px">
        <template slot-scope="scope">
          <svg-icon v-for="n in +scope.row.importance" :key="n" icon-class="star" class="meta-item__icon" />
        </template>
      </el-table-column>
      <el-table-column prop="about" label="概要" width="300">
      </el-table-column>
      <el-table-column prop="context" label="内容" width="200">
      </el-table-column>
      <el-table-column prop="status" label="状态" width="150">
        <el-tag type="success">已发布</el-tag>
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="handleClick(scope.row)" type="text" size="small">查看</el-button>
          <el-button @click="handleClick(scope.row)" type="text" size="small">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog title="" :visible.sync="dialogVisible" width="30%" :before-close="handleClose">
      <el-form :model="newNotice" label-width="80px" label-position="left">
        <el-form-item label="撰写人">
          <el-input v-model="newNotice.name" />
        </el-form-item>
        <el-form-item label="所属部门">
          <el-select class="name" v-model="newNotice.bumen" placeholder="请选择">
            <el-option v-for="item in docutypes" :key="item.value" :label="item.label" :value="item.label">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="newNotice.title" :autosize="{ minRows: 1, maxRows: 4 }" type="textarea" />
        </el-form-item>
        <el-form-item label="发布日期">
          <el-date-picker class="date" v-model="newNotice.date" type="date" placeholder="选择日期" value-format="timestamp">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="优先级">
          <el-rate v-model="newNotice.importance" :max="5" :colors="['#99A9BF', '#F7BA2A', '#FF9900']"
            :low-threshold="1" :high-threshold="5" style="display:inline-block" />
        </el-form-item>
        <el-form-item label="概要">
          <el-input v-model="newNotice.about" :autosize="{ minRows: 1, maxRows: 4 }" type="textarea" />
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="newNotice.context" :autosize="{ minRows: 4, maxRows: 4 }" type="textarea" />
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="quxiao">取消</el-button>
        <el-button type="primary" @click="sendNotice">提交</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const defaultNotice = {
  date: '',
  name: '',
  bumen: '',
  title: '',
  about: '',
  context: '',
  importance: 0
}
export default {
  data() {
    return {
      centerDialogVisible: false,
      dialogVisible: false,
      newNotice: defaultNotice,
      docutypes: [{
        value: '选项1',
        label: '市发展改革委'
      }, {
        value: '选项2',
        label: '市司法局'
      }, {
        value: '选项3',
        label: '市宣传部'
      }, {
        value: '选项4',
        label: '市组织部'
      }, {
        value: '选项5',
        label: '市发展改革委'
      }, {
        value: '选项6',
        label: '市人社局'
      }, {
        value: '选项7',
        label: '市气象局'
      }, {
        value: '选项8',
        label: '市政府办'
      }, {
        value: '选项9',
        label: '市交通运输局'
      }, {
        value: '选项10',
        label: '市卫健委'
      }, {
        value: '选项11',
        label: '市消防支援中队'
      }, {
        value: '选项12',
        label: '市市场监管局'
      }, {
        value: '选项13',
        label: '市公路局'
      }, {
        value: '选项14',
        label: '市统计局'
      }],
      tableData: [{
        date: '2023-01-26',
        name: '黄sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '已发布',
        context: '66666',
        importance: 4
      }, {
        date: '2023-01-26',
        name: '刘sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '草稿',
        context: '66666',
        importance: 3
      }, {
        date: '2023-01-26',
        name: '梁sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '已发布',
        context: '66666',
        importance: 5
      }, {
        date: '2023-01-26',
        name: '黄sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '草稿',
        context: '66666',
        importance: 2
      }, {
        date: '2023-01-26',
        name: '梁sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '已发布',
        context: '66666',
        importance: 5
      }, {
        date: '2023-01-26',
        name: '黄sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '草稿',
        context: '66666',
        importance: 2
      }, {
        date: '2023-01-26',
        name: '刘sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '草稿',
        context: '66666',
        importance: 3
      }, {
        date: '2023-01-26',
        name: '梁sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '已发布',
        context: '66666',
        importance: 5
      }, {
        date: '2023-01-26',
        name: '黄sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '草稿',
        context: '66666',
        importance: 2
      }, {
        date: '2023-01-26',
        name: '梁sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '已发布',
        context: '66666',
        importance: 5
      }, {
        date: '2023-01-26',
        name: '黄sir',
        bumen: '省长办公室',
        title: '关于单身男性的若干处理办法',
        about: '关于单身男性的若干处理办法',
        status: '草稿',
        context: '66666',
        importance: 2
      }]
    }
  },
  
methods: {
  resetNotice() {
    this.newNotice = {
      date: '',
      name: '',
      bumen: '',
      title: '',
      about: '',
      context: '',
      importance: 0
    }
  },
  handleClose(done) {
    this.$confirm('确认关闭？')
      .then(_ => {
        this.resetNotice()
        done();
      })
      .catch(_ => { });
  },
  handleClick(a) {
    this.dialogVisible = true
    this.newNotice = a
  },
  sendNotice() {
    this.dialogVisible = false
    console.log(this.newNotice)
    this.resetNotice()
  },
  quxiao() {
    this.dialogVisible = false
    this.resetNotice()
  }
},

}
</script>
<style scoped>

</style>
