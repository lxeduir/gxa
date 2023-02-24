<template>
  <div class="container">
    <el-table :data="tableData.admininfo" style="width: 100%" border>
      <el-table-column label="名称" width="250">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.Username }}</span>
        </template>
      </el-table-column>
      <el-table-column label="邮箱" width="250">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.Email }}</span>
        </template>
      </el-table-column>

      <el-table-column label="状态" width="250">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.State }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <!-- <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">编辑</el-button> -->
          <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog title="提示" :visible.sync="centerDialogVisible" width="20%" center>
      <span>编辑功能还没写！</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="centerDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="centerDialogVisible = false">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
import store from '@/store';
import { deleteUser } from '@/api/permisson';
import { Message } from 'element-ui';
import { getInfo } from '@/api/user';
export default {
  data() {
    return {
      tableData: store.getters.user,
      centerDialogVisible: false

    }
  },

  methods: {
    async refresh() {
      await store.dispatch('user/getInfo')
      this.tableData = store.getters.user
    },
    handleEdit(index, row) {
      this.centerDialogVisible = true
      console.log(index, row);
    },
    async handleDelete(index, row) {
      let res = await deleteUser(row.Uid)
      console.log(row.Uid)
      Message({
        center: true,
        message: res.data.msg,
        type: 'success',
        duration: 3 * 1000
      })
      this.refresh()
      console.log(index, row);
    },
    async fetchData() {
      let res = await getuserList()
    },
  }
}
</script>
<style>
.container {
  /* margin-top: 30px; */
}
</style>