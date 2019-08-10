<template>
    <div class="Tag">
        <content-item>

            <template slot="header">
                <a-select style="width: 200px;margin-right: 5px;" @select="reloadList" :defaultValue="0">
                    <a-select-option v-for="cate in category" :key="cate.Id" :value="cate.Id">{{cate.Title}}</a-select-option>
                </a-select>
                <a-button type="primary" @click="showHandleTag(false)" icon="plus">添加</a-button>
            </template>

            <div class="list" slot="body">
                <a-list itemLayout="horizontal" :dataSource="Tag" >
                    <a-list-item slot="renderItem" slot-scope="item,index">
                        <a-list-item-meta :description="item.Description">
                            <h4 slot="title">图板：{{item.BoardName}}</h4>
                            <h4 slot="title">简称：{{item.ShortName}}</h4>
                            <h4 slot="title">标题：{{item.Title}}</h4>
                        </a-list-item-meta>
                        <a slot="actions" @click="showHandleTag(true, item)">编辑</a>
                        <a slot="actions" @click="deleteTag(item.Id)">删除</a>
                    </a-list-item>
                </a-list>
            </div>
        </content-item>
        <a-modal title="添加标签" :visible="addTagVisible" :footer="null"  @cancel="()=> this.addTagVisible = false">
            <a-form >
                <a-form-item label="标签标题">
                    <a-input v-model="TagForm.Title" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="标签简称">
                    <a-input  v-model="TagForm.ShortName" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="选择分类">
                    <a-select v-model="TagForm.Cid">
                        <a-select-option v-for="cate in category" :key="cate.Id" :value="cate.Id">{{cate.Title}}</a-select-option>
                    </a-select>
                </a-form-item>
                <a-form-item label="标签描述">
                    <a-input v-model="TagForm.Description" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="图板名称">
                    <a-input  v-model="TagForm.BoardName" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" html-type="submit" class="login-form-button" :loading="addTagLoading" @click.prevent="handleTag">提交</a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form,Input, List,Select} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {listTagApi, addTagApi, deleteTagApi, updateTagApi} from "../../api/tag"
    import {listCategoryApi} from "../../api/category"
    export default {
        name: "Tag",
        components: {
            ContentItem,
            [List.name]: List,
            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
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
            this.listTag()
        },
        data() {
            return {
                addTagVisible: false,
                addTagLoading: false,
                Tag : [],
                category:[],
                activeCategory:0,
                DefaultTagForm:{
                    Id:0,
                    Cid:2,
                    Title:"",
                    ShortName:"",
                    Description:"",
                    BoardName:"",
                },
                TagForm:{
                    Id:0,
                    Cid:2,
                    Title:"",
                    ShortName:"",
                    Description:"",
                    BoardName:"",
                }
            }
        },
        methods:{
            showHandleTag(isUpdate, data){
                this.addTagVisible =true;
                if (!isUpdate) {
                    this.DefaultTagForm.Cid = this.activeCategory;
                    this.TagForm = Object.assign({},this.DefaultTagForm);
                    console.log("add",data);
                    return
                }
                console.log("update",data);
                this.TagForm = data
            },
            handleTag(){
                if (this.addTagLoading){
                    return
                }
                this.addTagLoading = true;
                if (this.TagForm.Id === 0) {
                    this.addTag()
                }else{
                    this.updateTag()
                }

            },
            listTag(){
                listTagApi({cid:this.activeCategory}).then(res => {
                    if (!res) {
                        return
                    }
                    this.Tag = res.data
                })
            },
            addTag(){
                addTagApi(this.createFormData()).then(res=>{
                    this.addTagLoading = false;
                    if (!res){
                        return
                    }
                    this.addTagVisible = false
                    this.listTag()
                })
            },
            updateTag(){
                updateTagApi(this.createFormData()).then(res=>{
                    this.addTagLoading = false;
                    if (!res){
                        return
                    }
                    this.addTagVisible = false
                    this.listTag()
                })
            },
            deleteTag(id){
                const form = new FormData;
                form.append("id", id);
                deleteTagApi(form).then(res=>{
                    if (!res){
                        return
                    }
                    const index = this.Tag.findIndex((c) => {
                        return c.Id === id;
                    });
                    this.Tag.splice(index, 1);
                })
            },
            createFormData(){
                const form = new FormData();
                form.append("id", this.TagForm.Id);
                form.append("cid", this.TagForm.Cid);
                form.append("title", this.TagForm.Title);
                form.append("shortName", this.TagForm.ShortName);
                form.append("description", this.TagForm.Description);
                form.append("boardName", this.TagForm.BoardName);
                return form
            },
            listCategory(){
                listCategoryApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.category = res.data;
                    this.category.splice(0,0,{
                        Id:0,
                        Title:"选择分类"
                    });
                })
            },
            reloadList(value){
                this.activeCategory = value;
                this.Tag = [];
                this.listTag()
            }
        }
    }
</script>

<style scoped>
    .add {
        margin: 0 -10px;
        border-bottom: 1px solid #ececec;
    }
</style>
