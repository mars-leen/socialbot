<template>
    <div class="category">
        <content-item>
            <a-button slot="header" type="primary" @click="showHandleCategory(false)" icon="plus">添加</a-button>
            <div class="list" slot="body">

                <a-list  itemLayout="vertical"  :dataSource="category"  :loading="loading">
                    <a-list-item  slot="renderItem" slot-scope="item,index">
                        <a-list-item-meta :description="item.Description">
                            <h4 slot="title">简称：{{item.ShortName}} </h4>
                            <h4 slot="title">标题：{{item.Title}}</h4>
                        </a-list-item-meta>
                        <img v-if="item.Cover" slot="extra"  alt="cover" :src="item.Cover" />
                        <a slot="actions" @click="showHandleCategory(true, item)">编辑</a>
                        <a slot="actions" @click="deleteCategory(item.Id)">删除</a>
                    </a-list-item>
                </a-list>

            </div>
        </content-item>
        <a-modal title="添加分类" :visible="addCategoryVisible" :footer="null"  @cancel="()=> this.addCategoryVisible = false">
            <a-form id="components-form-demo-normal-login" >
                <a-form-item  label="分类标题">
                    <a-input v-model="categoryForm.Title" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="分类简称">
                    <a-input v-model="categoryForm.ShortName" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="分类描述">
                    <a-input v-model="categoryForm.Description" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="分类排序">
                    <a-input  v-model="categoryForm.Sort" type="integer" >
                    </a-input>
                </a-form-item>
                <a-form-item label="封面">
                    <div class="cover" v-if="categoryForm.Cover">{{categoryForm.Cover}}</div>
                    <a-upload :fileList="fileList" :beforeUpload ="handleCover" :remove="handleRemove" :multiple="false">
                        <a-button v-if="fileList.length<1"><a-icon type="upload" />
                            <span v-if="categoryForm.Cover">替换</span>
                            <span v-else>上传</span>
                        </a-button>
                    </a-upload>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" html-type="submit" class="login-form-button" :loading="addCategoryLoading" @click.prevent="handleCategory">提交</a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form,Input, List, Upload} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {listCategoryApi, addCategoryApi, deleteCategoryApi, updateCategoryApi} from "../../api/category"
    export default {
        name: "Category",
        components: {
            ContentItem,
            [Upload.name]: Upload,
            [List.name]: List,
            [List.Item.name]: List.Item,
            [List.Item.Meta.name]: List.Item.Meta,
            [Button.name]: Button,
            [Icon.name]: Icon,
            [Modal.name]: Modal,
            [Form.name]: Form,
            [Form.Item.name]: Form.Item,
            [Input.name]: Input,
        },
        created(){
            this.listCategory()
        },
        data() {
            return {
                loading:false,
                addCategoryVisible: false,
                addCategoryLoading: false,
                category : [],
                fileList:[],
                defaultCategoryForm:{
                    Id:0,
                    Title:"",
                    ShortName:"",
                    Description:"",
                    Sort:0,
                    Cover:"",
                },
                categoryForm:{
                    Id:0,
                    Title:"",
                    ShortName:"",
                    Description:"",
                    Sort:0,
                    Cover:"",
                }
            }
        },
        methods:{
            showHandleCategory(isUpdate, data){
                this.addCategoryVisible =true;
                if (!isUpdate) {
                    this.categoryForm = Object.assign({},this.defaultCategoryForm);
                    this.fileList = [];
                    return
                }
                this.categoryForm = data
            },
            handleCategory(){
                if (this.addCategoryLoading){
                    return
                }
                this.addCategoryLoading = true;
                if (this.categoryForm.Id === 0) {
                    this.addCategory()
                }else{
                    this.updateCategory()
                }
            },
            listCategory(){
                this.loading=true;
                listCategoryApi().then(res => {
                    this.loading=false;
                    if (!res) {
                        return
                    }
                    this.category = res.data
                })
            },
            addCategory(){
                addCategoryApi(this.createFormData()).then(res=>{
                    this.addCategoryLoading = false;
                    if (!res){
                        return
                    }
                    this.addCategoryVisible = false;
                    this.listCategory();
                })
            },
            updateCategory(){
                updateCategoryApi(this.createFormData()).then(res=>{
                    this.addCategoryLoading = false;
                    if (!res){
                        return
                    }
                    this.addCategoryVisible = false;
                    this.listCategory();
                })
            },
            deleteCategory(id){
                const form = new FormData;
                form.append("id", id);
                deleteCategoryApi(form).then(res=>{
                    if (!res){
                        return
                    }
                    const index = this.category.findIndex((c) => {
                        return c.Id === id;
                    });
                    this.category.splice(index, 1);
                })
            },
            createFormData(){
                const form = new FormData();
                form.append("id", this.categoryForm.Id);
                form.append("title", this.categoryForm.Title);
                form.append("shortName", this.categoryForm.ShortName);
                form.append("description", this.categoryForm.Description);
                form.append("sort", this.categoryForm.Sort);
                if (this.fileList.length >0){
                    form.append("cover", this.fileList[0]);
                }
                return form
            },
            handleCover(file){
                this.fileList = [...this.fileList, file]
                return false;
            },
            handleRemove(file) {
                const index = this.fileList.indexOf(file);
                const newFileList = this.fileList.slice();
                newFileList.splice(index, 1);
                this.fileList = newFileList
            },
        }
    }
</script>

<style scoped>
</style>
