<template>
    <div class="Copywriter">
        <content-item>
            <a-button slot="header" type="primary" @click="showHandleCopywriter(false)" icon="plus">添加</a-button>
            <div class="list" slot="body">
                <a-list itemLayout="horizontal" :dataSource="Copywriter" >
                    <a-list-item slot="renderItem" slot-scope="item,index">
                        <a-list-item-meta :description="item.Description">
                            <h4 slot="title">{{item.Title}}</h4>
                        </a-list-item-meta>
                        <a slot="actions" @click="showHandleCopywriter(true, item)">编辑</a>
                        <a slot="actions" @click="deleteCopywriter(item.Id)">删除</a>
                    </a-list-item>
                </a-list>
            </div>
        </content-item>
        <a-modal title="添加分类" :visible="addCopywriterVisible" :footer="null"  @cancel="()=> this.addCopywriterVisible = false">
            <a-form id="components-form-demo-normal-login" >
                <a-form-item  label="标题">
                    <a-input v-model="CopywriterForm.Title" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="分类描述">
                    <a-input v-model="CopywriterForm.Description" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" html-type="submit" class="login-form-button" :loading="addCopywriterLoading" @click.prevent="handleCopywriter">提交</a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form,Input, List} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {listCopywriterApi, addCopywriterApi, deleteCopywriterApi, updateCopywriterApi} from "../../api/copywriter"
    export default {
        name: "Copywriter",
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
            this.listCopywriter()
        },
        data() {
            return {
                addCopywriterVisible: false,
                addCopywriterLoading: false,
                Copywriter : [],
                defaultCopywriterForm:{
                    Id:0,
                    Title:"",
                    Description:"",
                },
                CopywriterForm:{
                    Id:0,
                    Title:"",
                    Description:"",
                }
            }
        },
        methods:{
            showHandleCopywriter(isUpdate, data){
                this.addCopywriterVisible =true;
                if (!isUpdate) {
                    this.CopywriterForm = Object.assign({},this.defaultCopywriterForm);
                    return
                }
                this.CopywriterForm = data
            },
            handleCopywriter(){
                if (this.addCopywriterLoading){
                    return
                }
                this.addCopywriterLoading = true;
                if (this.CopywriterForm.Id === 0) {
                    this.addCopywriter()
                }else{
                    this.updateCopywriter()
                }
            },
            listCopywriter(){
                listCopywriterApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.Copywriter = res.data
                })
            },
            addCopywriter(){
                addCopywriterApi(this.createFormData()).then(res=>{
                    this.addCopywriterLoading = false;
                    if (!res){
                        return
                    }
                    this.addCopywriterVisible = false;
                    this.Copywriter.push(this.CopywriterForm)
                })
            },
            updateCopywriter(){
                updateCopywriterApi(this.createFormData()).then(res=>{
                    this.addCopywriterLoading = false;
                    if (!res){
                        return
                    }
                    this.addCopywriterVisible = false
                })
            },
            deleteCopywriter(id){
                const form = new FormData;
                form.append("id", id);
                deleteCopywriterApi(form).then(res=>{
                    if (!res){
                        return
                    }
                    const index = this.Copywriter.findIndex((c) => {
                        return c.Id === id;
                    });
                    this.Copywriter.splice(index, 1);
                })
            },
            createFormData(){
                const form = new FormData();
                form.append("id", this.CopywriterForm.Id);
                form.append("title", this.CopywriterForm.Title);
                form.append("description", this.CopywriterForm.Description);
                return form
            }
        }
    }
</script>

<style scoped>
</style>
