<template>
    <div class="category">
        <content-item>
            <a-button slot="header" type="primary" @click="showHandleCategory(false)" icon="plus">添加</a-button>
            <div class="list" slot="body">
                <a-list itemLayout="horizontal" :dataSource="category" >
                    <a-list-item slot="renderItem" slot-scope="item,index">
                        <a-list-item-meta :description="item.Description">
                            <h4 slot="title">{{item.Title}}({{item.ShortName}})</h4>
                        </a-list-item-meta>
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
                <a-form-item>
                    <a-button type="primary" html-type="submit" class="login-form-button" :loading="addCategoryLoading" @click.prevent="handleCategory">提交</a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form,Input, List} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {listCategoryApi, addCategoryApi, deleteCategoryApi, updateCategoryApi} from "../../api/category"
    export default {
        name: "Category",
        components: {
            ContentItem,
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
                addCategoryVisible: false,
                addCategoryLoading: false,
                category : [],
                defaultCategoryForm:{
                    Id:0,
                    Title:"",
                    ShortName:"",
                    Description:"",
                    Sort:0,
                },
                categoryForm:{
                    Id:0,
                    Title:"",
                    ShortName:"",
                    Description:"",
                    Sort:0,
                }
            }
        },
        methods:{
            showHandleCategory(isUpdate, data){
                this.addCategoryVisible =true;
                if (!isUpdate) {
                    this.categoryForm = Object.assign({},this.defaultCategoryForm);
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
                listCategoryApi().then(res => {
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
                    this.addCategoryVisible = false
                    this.category.push(this.categoryForm)
                })
            },
            updateCategory(){
                updateCategoryApi(this.createFormData()).then(res=>{
                    this.addCategoryLoading = false;
                    if (!res){
                        return
                    }
                    this.addCategoryVisible = false
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
                form.append("sort", this.categoryForm.Sort)
                return form
            }
        }
    }
</script>

<style scoped>
</style>
