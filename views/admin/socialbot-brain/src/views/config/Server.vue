<template>
    <div class="Server">
        <content-item>
            <a-button slot="header" type="primary" @click="showHandleServer(false)" icon="plus">添加</a-button>
            <div slot="body">
                <a-row :gutter="16">
                    <a-col v-for="s in server" :key="s.Id" :sm="1" :md="8" class="server-card">
                        <a-card  :title="s.Title">
                            <template class="ant-card-actions" slot="actions">
                                <a-icon type="edit" @click="showHandleServer(true, s)"/>
                                <a-icon type="delete" @click="deleteServer(s.Id)" />
                            </template>
                            <p style="word-wrap: break-word"><strong>ApiKey</strong>: {{s.ApiKey}}</p>
                            <div><strong>RegionID</strong>: {{s.RegionID}}</div>
                            <div><strong>PlanId</strong>: {{s.PlanId}}</div>
                            <div><strong>OsId</strong>: {{s.OsId}}</div>
                            <div><strong>scriptName</strong>: {{s.ScriptName}}</div>
                            <div><strong>scriptType</strong>: {{s.ScriptType}}</div>
                        </a-card>
                    </a-col>
                </a-row>
            </div>
        </content-item>
        <a-modal title="添加分类" :visible="addServerVisible" :footer="null" @cancel="()=> this.addServerVisible = false">
            <a-form>
                <a-form-item   label="标题">
                    <a-input v-model="serverForm.Title" type="string">
                    </a-input>
                </a-form-item>
                <a-form-item  label="apiKey">
                    <a-input v-model="serverForm.ApiKey" type="string">
                    </a-input>
                </a-form-item>
                <a-form-item  label="地区id">
                    <a-input v-model="serverForm.RegionID" type="string">
                    </a-input>
                </a-form-item>
                <a-form-item  label="计划id">
                    <a-input v-model="serverForm.PlanId" type="integer">
                    </a-input>
                </a-form-item>
                <a-form-item  label="系统id">
                    <a-input v-model="serverForm.OsId" type="string">
                    </a-input>
                </a-form-item>

                <a-form-item  label="startScript名称">
                    <a-input v-model="serverForm.ScriptName" type="string">
                    </a-input>
                </a-form-item>
                <a-form-item  label="startScript类型">
                    <a-input v-model="serverForm.ScriptType" type="integer">
                    </a-input>
                </a-form-item>
                <a-form-item  label="startScript内容">
                    <a-textarea v-model="serverForm.ScriptContent" placeholder="script content"
                                :autosize="{ minRows: 6, maxRows: 16 }"></a-textarea>
                </a-form-item>
                <a-form-item  >
                    <a-button type="primary" html-type="submit" class="login-form-button" :loading="addServerLoading"
                              @click.prevent="handleServer">提交
                    </a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form, Input, List,Card, Row, Col} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {
        listServerConfigApi,
        addServerConfigApi,
        deleteServerConfigApi,
        updateServerConfigApi
    } from "../../api/server"

    export default {
        name: "ConfigServer",
        components: {
            ContentItem,
            [List.name]: List,
            [Card.name]: Card,
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
            this.listServer()
        },
        data() {
            return {
                addServerVisible: false,
                addServerLoading: false,
                server: [],
                defaultServerForm: {
                    Id: 0,
                    Title: "server",
                    ScriptName: "socailbot_script",
                    ScriptType: "boot",
                    ScriptContent: "#!/bin/sh\n" +
                        "\n" +
                        "\n" +
                        "# NOTE: This is an example that sets up SSH authorization. To use it, you'd need to replace \"ssh-rsa AA... youremail@example.com\" with your SSH public.\n" +
                        "# You can replace this entire script with anything you'd like, there is no need to keep it\n" +
                        "\n" +
                        "mkdir /root/brook\n" +
                        "cd /root/brook\n" +
                        "wget https://github.com/txthinking/brook/releases/download/v20190601/brook\n" +
                        "chmod 777 brook\n" +
                        "firewall-cmd --permanent --zone=public --add-port=8888/tcp\n" +
                        "firewall-cmd --reload\n" +
                        "\n" +
                        "echo '1' > blocked\n" +
                        "\n" +
                        "nohup ./brook -d server -l :8888 -p 543284@Lmq >> run.log 2>&1 \n" +
                        "&\n" +
                        "\n" +
                        "echo '2' > blocked",
                    ApiKey: "",
                    RegionID: 1,
                    PlanId: 200,
                    OsId: 167
                },
                serverForm: {
                    Id: 0,
                    Title: "server",
                    ScriptName: "socailbot_script",
                    ScriptType: "boot",
                    ScriptContent: "#!/bin/sh\n" +
                        "\n" +
                        "\n" +
                        "# NOTE: This is an example that sets up SSH authorization. To use it, you'd need to replace \"ssh-rsa AA... youremail@example.com\" with your SSH public.\n" +
                        "# You can replace this entire script with anything you'd like, there is no need to keep it\n" +
                        "\n" +
                        "mkdir /root/brook\n" +
                        "cd /root/brook\n" +
                        "wget https://github.com/txthinking/brook/releases/download/v20190601/brook\n" +
                        "chmod 777 brook\n" +
                        "firewall-cmd --permanent --zone=public --add-port=8888/tcp\n" +
                        "firewall-cmd --reload\n" +
                        "\n" +
                        "echo '1' > blocked\n" +
                        "\n" +
                        "nohup ./brook -d server -l :8888 -p 543284@Lmq >> run.log 2>&1 \n" +
                        "&\n" +
                        "\n" +
                        "echo '2' > blocked",
                    ApiKey: "",
                    RegionID: 1,
                    PlanId: 200,
                    OsId: 167
                }
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
                this.addServerLoading = false;
                if (this.serverForm.Id === 0) {
                    this.addServer()
                } else {
                    this.updateServer()
                }
            },
            listServer() {
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
                    this.server = list
                })
            },
            addServer() {
                addServerConfigApi(this.createFormData()).then(res => {
                    this.addServerLoading = false;
                    if (!res) {
                        return
                    }
                    this.addServerVisible = false;
                    this.listServer()
                })
            },
            updateServer() {
                updateServerConfigApi(this.createFormData()).then(res => {
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
                deleteServerConfigApi(form).then(res => {
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
                form.append("title", this.serverForm.Title);
                form.append("value", JSON.stringify(this.serverForm));
                return form
            }
        }
    }
</script>

<style scoped>
    .server-card {
        margin-bottom: 16px;
    }
</style>
