<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="URL">
          <el-input v-model="searchInfo.url" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="客户账号">
          <el-input v-model="searchInfo.username" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="el-icon-delete"
              size="mini"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
            >删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="URL" prop="url" width="" />
        <el-table-column align="left" label="用户名" prop="username" width="120" />
        <el-table-column align="left" label="密码" prop="password" width="120" />
        <el-table-column align="left" label="客户ID" prop="user_id" width="120" />
        <el-table-column align="left" label="客户名称" prop="user.nickName" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              icon="el-icon-edit"
              size="small"
              class="table-button"
              @click="updateJunquan(scope.row)"
            >变更
            </el-button>
            <el-button type="text" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="URL:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名:">
          <el-input v-model="formData.username" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:">
          <el-input v-model="formData.password" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="客户" prop="user_id">
          <el-select v-model="formData.user_id" placeholder="请选择客户" clearable :style="{width: '100%'}">
            <el-option
              v-for="(item, index) in users"
              :key="index"
              :label="item.nickName"
              :value="item.ID"
              :disabled="item.disabled"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createJunquan,
  deleteJunquan,
  deleteJunquanByIds,
  findJunquan,
  getJunquanList,
  updateJunquan
} from '@/api/junquan' //  此处请自行替换地址
import { getKehuList } from '@/api/user'
import infoList from '@/mixins/infoList'

export default {
  name: 'Junquan',
  mixins: [infoList],
  data() {
    return {
      searchInfo:{},
      users: [],
      getKehuApi:getKehuList,
      listApi: getJunquanList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        url: 'http://supergj.junquan.com.cn',
        username: '',
        password: '8d59012a1c6251fcea67331d1352370a7b739d6b',
        user_id: '',
      }
    }
  },
  async created() {
    await this.getTableData()
    this.getKehuApi({ page: 1, pageSize: 10000 }).then(res => {
      console.log(res.data.list)

      this.users = res.data.list
    })
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteJunquan(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteJunquanByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateJunquan(row) {
      const res = await findJunquan({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reJq
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        url: '',
        username: '',
        password: '8d59012a1c6251fcea67331d1352370a7b739d6b',
        user_id: '',
      }
    },
    async deleteJunquan(row) {
      const res = await deleteJunquan({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createJunquan(this.formData)
          break
        case 'update':
          res = await updateJunquan(this.formData)
          break
        default:
          res = await createJunquan(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
