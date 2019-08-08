<template>
    <div class="Server">
        <content-item>
            <a-button slot="header" type="primary" @click="showHandleServer(false)" icon="plus">添加</a-button>
            <div slot="body">
                <a-row :gutter="16">
                    <a-col v-for="s in server" :key="s.Id" :sm="1" :md="8" class="server-card">
                        <a-card  :title="s.label">
                            <a-spin :spinning="(s.server_state !=='ok')" :delay="500" :tip="'服务器部署中:'+ s.server_state ">
                                <div><strong>创建日期</strong>: {{s.date_created}}</div>
                                <div><strong>总的流量</strong>: {{s.allowed_bandwidth_gb}}</div>
                                <div><strong>已用流量</strong>: {{s.current_bandwidth_gb}}</div>
                                <div><strong>已用金额</strong>: {{s.pending_charges}}</div>
                                <div><strong>默认密码</strong>: {{s.default_password}}</div>
                                <div><strong>地理位置</strong>: {{s.location}}</div>
                                <div><strong>外网地址</strong>: {{s.main_ip}}</div>
                                <div><strong>server_state</strong>: {{s.server_state}}</div>
                                <div><strong>status</strong>: {{s.status}}</div>
                                <div><strong>power_status</strong>: {{s.power_status}}</div>
                            </a-spin>
                            <template class="ant-card-actions" slot="actions">
                                <a-icon type="delete" @click="deleteServer(s.Id)" />
                            </template>
                        </a-card>
                    </a-col>
                </a-row>
            </div>
        </content-item>
        <a-modal title="添加" :visible="addServerVisible" :footer="null" @cancel="()=> this.addServerVisible = false">
            <a-form>
                <a-form-item label="名称">
                    <a-input v-model="serverForm.Name" type="string">
                    </a-input>
                </a-form-item>
                <a-form-item  label="选择服务器配置">
                    <a-select  @select="changeConfig">
                        <a-select-option v-for="c in serverConfig" :key="c.Id" :value="c.Id">{{c.Title}}</a-select-option>
                    </a-select>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" html-type="submit" class="login-form-button" :loading="addServerLoading" @click.prevent="handleServer">提交
                    </a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form, Input, List, Card, Row, Col, Select,Tag, Spin} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {
        listRobotServerApi,
        addRobotServerApi,
        deleteRobotServerApi,
    } from "../../api/robotServer"
    import {
        listServerConfigApi,
    } from "../../api/server"

    export default {
        name: "Server",
        components: {
            ContentItem,
            [List.name]: List,
            [Spin.name]: Spin,
            [Tag.name]: Tag,
            [Card.name]: Card,
            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
            [Row.name]: Row,
            [Col.name]: Col,
            [Card.Meta.name]: Card.Meta,
            [List.Item.name]: List.Item,
            [List.Item.Meta.name]: List.Item.Meta,
            [Button.name]: Button,
            [Icon.name]: Icon,
            [Modal.name]: Modal,
            [Form.name]: Form,
            [Form.Item.name]: Form.Item,
            [Input.name]: Input,
            [Input.TextArea.name]: Input.TextArea,
        },
        created() {
            this.listServer();
            this.listServerConfig();
        },
        data() {
            return {
                date:"",
                addServerVisible: false,
                addServerLoading: false,
                server: [],
                serverConfig: [],
                defaultServerForm: {
                    Id: 0,
                    Name: "",
                    Cid: 0,
                },
                serverForm: {
                    Id: 0,
                    Name: "",
                    Cid: 0,
                }
            }
        },
        mounted() {
            let _this = this;
            this.timer = setInterval(() => {
                _this.listServer()
            }, 3000)
        },
        beforeDestroy() {
            if (this.timer) {
                clearInterval(this.timer);
            }
        },
        methods: {
            showHandleServer(isUpdate, data) {
                this.addServerVisible = true;
                if (!isUpdate) {
                    this.serverForm = Object.assign({}, this.defaultServerForm);
                    return
                }
                this.serverForm = data
            },
            handleServer() {
                if (this.addServerLoading) {
                    return
                }
                this.addServerLoading = true;
                if (this.serverForm.Id === 0) {
                    this.addServer()
                }
            },
            listServer() {
                listRobotServerApi().then(res => {
                    if (!res) {
                        return
                    }
                    const result = res.data;
                    const list = [];
                    for (let i in result) {
                        let s = JSON.parse(result[i].FullServerInfo);
                        s.Id = result[i].Id;
                        list.push(s)
                    }
                    this.server = list;
                })
            },
            addServer() {
                addRobotServerApi(this.createFormData()).then(res => {
                    this.addServerLoading = false;
                    if (!res) {
                        return
                    }
                    this.addServerVisible = false;
                    this.listServer()
                })
            },
            deleteServer(id) {
                const form = new FormData;
                form.append("id", id);
                deleteRobotServerApi(form).then(res => {
                    if (!res) {
                        return
                    }
                    const index = this.server.findIndex((c) => {
                        return c.Id === id;
                    });
                    this.server.splice(index, 1);
                })
            },
            createFormData() {
                const form = new FormData();
                form.append("id", this.serverForm.Id);
                form.append("name", this.serverForm.Name);
                form.append("cid", this.serverForm.Cid);
                return form
            },
            listServerConfig() {
                listServerConfigApi().then(res => {
                    if (!res) {
                        return
                    }
                    const result = res.data;
                    const list = [];
                    for (let i in result) {
                        let s = JSON.parse(result[i].Value);
                        s.Id = result[i].Id;
                        list.push(s)
                    }
                    this.serverConfig = list
                })
            },
            changeConfig(value){
                this.serverForm.Cid = value
            },
        }
    }
</script>

<style scoped>
    .server-card {
        margin-bottom: 16px;
    }
</style>
