<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddRole">新增角色</el-button>
    <!-- "CreateTime":"2022-12-03T15:09:13+08:00","UpdateBy":"等待实现","UpdateTime":"2022-12-03T15:09:13+08:00","Orders":1},{"RoleId":"2","Name":"guest","PermissionJson":{"list":[{"type":"0","type_id":"0","permission":"0"}]},"Description":"访客","CreateBy":"等待实现","CreateTime":"2022-12-03T15:22:30+08:00","UpdateBy":"等待实现","UpdateTime":"2022-12-03T15:22:30+08:00","Orders":2},{"RoleId":"3","Name":"测试","PermissionJson":{"list":[{"type":"rout","type_id":"11","permission":"1"},{"type":"rout","type_id":"16","permission":"1"},{"type":"rout","type_id":"17","permission":"1"},{"type":"rout","type_id":"21","permission":"1"},{"type":"rout","type_id":"26","permission":"1"},{"type":"rout","type_id":"24","permission":"1"}]},"Description":"测试","CreateBy":"等待实现","CreateTime":"2023-01-12T22:01:24+08:00","UpdateBy":"等待实现","UpdateTime": -->
    <el-table :data="rolesList" style="width: 100%;margin-top:30px;" border>
      <el-table-column align="center" label="RoleId" width="220">
        <template slot-scope="scope">
          {{ scope.row.RoleId }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="角色名称" width="220">
        <template slot-scope="scope">
          {{ scope.row.Name }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="权限描述">
        <template slot-scope="scope">
          {{ scope.row.Description }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="创建时间">
        <template slot-scope="scope">
          {{ scope.row.CreateTime }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="修改时间">
        <template slot-scope="scope">
          {{ scope.row.UpdateTime }}
        </template>
      </el-table-column>

      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button type="primary" size="small" @click="handleEdit(scope)">修改</el-button>
          <el-button type="danger" size="small" @click="handleDelete(scope)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!--  点击新增角色时弹出的对话框  -->
    <el-dialog :visible.sync="dialogVisible" :title="dialogType === 'edit' ? 'Edit Role' : 'New Role'"
      :before-close="handleClose">
      <el-form :model="role" label-width="80px" label-position="left">

        <el-form-item label="角色名称">
          <el-input v-model="role.Name" />
        </el-form-item>
        <el-form-item label="权限描述">
          <el-input v-model="role.Description" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" />
        </el-form-item>

        <el-form-item label="Menus">
          <el-tree ref="tree" :check-strictly="checkStrictly" :data="routesData" :props="defaultProps" show-checkbox
            node-key="role_id" class="permission-tree" default-expand-all=true />
        </el-form-item>

      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmRole">提交</el-button>
      </div>

    </el-dialog>
    
  </div>
</template>

<script>
import path from 'path'
import { deepClone } from '@/utils'
import { role_search, role_add, role_del, role_update, path_search } from '@/api/permisson'
import { Message } from 'element-ui'

const defaultRole = {
  RoleId: '',
  Name: '',
  PermissionJson: {
    list: []
  },
  Description: '',
  Orders: null,

}
export default {
  data() {
    return {
      role: Object.assign({}, defaultRole),
      routes: [],
      rolesList: [],
      dialogVisible: false,
      dialogType: 'new',
      checkStrictly: true,
      defaultProps: {
        children: 'children',
        label: 'title'
      }
    }
  },
  computed: {
    routesData() {
      return this.routes
    }
  },
  created() {
    this.getRoutes()
    this.getRoles()
  },
  methods: {
    handleClose(done) {
      this.$confirm('确认关闭？')
        .then(_ => {
          done();
        })
        .catch(_ => { });
    },
    async getRoutes() {
      const res = await path_search()
      this.serviceRoutes = res.data.list
      this.routes = this.generateRoutes(res.data.list)
    },
    async getRoles() {
      const res = await role_search()
      this.rolesList = res.data.list
    },
    generateRoutes(routes, basePath = '/') {
      const res = []
      for (let route of routes) {
        //如果不是只有一个孩子要展示 就返回false
        const onlyOneShowingChild = this.onlyOneShowingChild(route.children, route)
        if (route.children && onlyOneShowingChild) {
          route = onlyOneShowingChild
        }
        const newRole = {
          type_id: String(route.role_id),
          type: 'rout',
          permission: '1'
        }
        const data = {
          path: path.resolve(basePath, route.path),
          title: route.meta && route.meta.title,
          role_id: newRole
        }
        // recursive child routes
        if (route.children) {
          data.children = this.generateRoutes(route.children, data.path)
        }
        res.push(data)
      }
      return res
    },
    handleAddRole() {
      if (this.routes.length == 0) {
        Message({
          center: true,
          message: '    对不起!您无权进行此操作!',
          type: 'error',
          duration: 5 * 1000
        })
        return
      }
      //浅复制一个对象
      this.role = Object.assign({}, defaultRole)
      if (this.$refs.tree) {
        this.$refs.tree.setCheckedNodes([])
      }
      this.dialogType = 'new'
      this.dialogVisible = true
    },
    handleEdit(scope) {
      if (this.routes.length == 0) {
        Message({
          center: true,
          message: '对不起!您无权进行此操作！',
          type: 'error',
          duration: 5 * 1000
        })
        return
      }
      this.dialogType = 'edit'
      this.dialogVisible = true
      this.checkStrictly = true
      // console.log(scope.row)
      this.role = deepClone(scope.row)
      this.$nextTick(() => {
        const routes = this.generateRoutes(this.role.routes)
        this.$refs.tree.setCheckedNodes(this.generateArr(routes))
        this.checkStrictly = false
      })
    },
    handleDelete({ $index, row }) {
      if (this.routes.length == 0) {
        Message({
          center: true,
          message: '对不起!您无权进行此操作！',
          type: 'error',
          duration: 5 * 1000
        })
        return
      }
      this.$confirm('确认移除这个角色吗', 'Warning', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          await role_del({ RoleId: row.RoleId })
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
          this.getRoles()
        })
        .catch(err => { console.error(err) })
    },
    generateTree(routes, basePath = '/', checkedKeys) {
      const res = []

      for (const route of routes) {
        const routePath = path.resolve(basePath, route.path)

        // recursive child routes
        if (route.children) {
          route.children = this.generateTree(route.children, routePath, checkedKeys)
        }

        if (checkedKeys.includes(routePath) || (route.children && route.children.length >= 1)) {
          res.push(route)
        }
      }
      return res
    },
    async confirmRole() {
      try {
        const isEdit = this.dialogType === 'edit'

        const checkedKeys = this.$refs.tree.getCheckedKeys()
        console.log(checkedKeys)
        if (!isEdit) {
          this.role.PermissionJson.list = checkedKeys
          await role_add(this.role)
          this.$message({
            type: 'success',
            message: '添加成功!'
          })
          this.role = Object.assign({}, defaultRole)
          this.getRoles()
          this.dialogVisible = false
        }
        else {
          // const upRole={
          //   RoleId:
          // }
        }
      }
      catch (error) {
        console.log(error)
      }
      // this.role.routes = this.generateTree(deepClone(this.serviceRoutes), '/', checkedKeys)

      // if (isEdit) {
      //   await updateRole(this.role.key, this.role)
      //   for (let index = 0; index < this.rolesList.length; index++) {
      //     if (this.rolesList[index].key === this.role.key) {
      //       this.rolesList.splice(index, 1, Object.assign({}, this.role))
      //       break
      //     }
      //   }
      // } else {
      //   const { data } = await addRole(this.role)
      //   this.role.key = data.key
      //   this.rolesList.push(this.role)
      // }

      // const { description, key, name } = this.role
      // this.dialogVisible = false
      // this.$notify({
      //   title: 'Success',
      //   dangerouslyUseHTMLString: true,
      //   message: `
      //       <div>Role Key: ${key}</div>
      //       <div>Role Name: ${name}</div>
      //       <div>Description: ${description}</div>
      //     `,
      //   type: 'success'
      // })
    },
    onlyOneShowingChild(showingChildren = [], parent) {
      // console.log("&&&&&&&")
      // console.log(showingChildren)
      // console.log(parent)
      let onlyOneChild = null
      // const showingChildren = children.filter(item => !item.hidden)

      // When there is only one child route, the child route is displayed by default
      if (showingChildren.length === 1) {
        onlyOneChild = showingChildren[0]
        onlyOneChild.path = path.resolve(parent.path, onlyOneChild.path)
        return onlyOneChild
      }

      // Show parent if there are no child route to display
      if (showingChildren.length === 0) {
        onlyOneChild = { ...parent, path: '', noShowingChildren: true }
        return onlyOneChild
      }

      return false
    }
  }
}
</script>

<style lang="scss" scoped>
.app-container {
  .roles-table {
    margin-top: 30px;
  }

  .permission-tree {
    margin-bottom: 30px;
  }
}
</style>
