<template>
    <div class="Server">
        <content-item>
            <a-button slot="header" type="primary" @click="showHandleServer(false)" icon="plus">添加</a-button>
        </content-item>
        <a-modal title="添加分类" :visible="addServerVisible" :footer="null"  @cancel="()=> this.addServerVisible = false">
            <a-form id="components-form-demo-normal-login" >
                <a-form-item label="分类标题">
                    <a-input v-model="ServerForm.Title" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="分类简称">
                    <a-input v-model="ServerForm.ShortName" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="分类描述">
                    <a-input v-model="ServerForm.Description" type="string" >
                    </a-input>
                </a-form-item>
                <a-form-item label="分类排序">
                    <a-input  v-model="ServerForm.Sort" type="integer" >
                    </a-input>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" html-type="submit" class="login-form-button" :loading="addServerLoading" @click.prevent="handleServer">提交</a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form,Input, List} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {listServerApi, addServerApi, deleteServerApi, updateServerApi} from "../../api/se"
    export default {
        name: "Server",
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
            this.listServer()
        },
        data() {
            return {
                addServerVisible: false,
                addServerLoading: false,
                Server : [],
                defaultServerForm:{
                    Id:0,
                    Title:"",
                    ScriptName:"",
                    ScriptType:"",
                    ScriptContent:"",
                    ApiKey:"",
                    RegionID:0,
                    PlanId:0,
                    OsId:0
                },
                ServerForm:{
                    Id:0,
                    Title:"",
                    ScriptName:"",
                    ScriptType:"",
                    ScriptContent:"",
                    ApiKey:"",
                    RegionID:0,
                    PlanId:0,
                    OsId:0
                }
            }
        },
        methods:{
            showHandleServer(isUpdate, data){
                this.addServerVisible =true;
                if (!isUpdate) {
                    this.ServerForm = Object.assign({},this.defaultServerForm);
                    return
                }
                this.ServerForm = data
            },
            handleServer(){
                if (this.addServerLoading){
                    return
                }
                this.addServerLoading = false;
                if (this.ServerForm.Id === 0) {
                    this.addServer()
                }else{
                    this.updateServer()
                }
            },
            listServer(){
                listServerApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.Server = res.data
                })
            },
            addServer(){
                addServerApi(this.createFormData()).then(res=>{
                    this.addServerLoading = false;
                    if (!res){
                        return
                    }
                    this.addServerVisible = false
                    this.Server.push(this.ServerForm)
                })
            },
            updateServer(){
                updateServerApi(this.createFormData()).then(res=>{
                    this.addServerLoading = false;
                    if (!res){
                        return
                    }
                    this.addServerVisible = false
                })
            },
            deleteServer(id){
                const form = new FormData;
                form.append("id", id);
                deleteServerApi(form).then(res=>{
                    if (!res){
                        return
                    }
                    const index = this.Server.findIndex((c) => {
                        return c.Id === id;
                    });
                    this.Server.splice(index, 1);
                })
            },
            createFormData(){
                const form = new FormData();
                form.append("id", this.ServerForm.Id);
                form.append("title", this.ServerForm.Title);
                form.append("shortName", this.ServerForm.ShortName);
                form.append("description", this.ServerForm.Description);
                form.append("sort", this.ServerForm.Sort)
                return form
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
