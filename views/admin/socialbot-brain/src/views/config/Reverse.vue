<template>
    <div class="ReverseHost">
        <content-item>
            <a-button slot="header" type="primary" @click="showHandleReverseHost(false)" icon="plus">添加</a-button>
            <div slot="body">
                <a-row :gutter="16">
                    <a-col v-for="s in ReverseHost" :key="s.Id" :sm="1" :md="12" class="ReverseHost-card">
                        <a-spin :spinning="s.Loading">
                            <a-card :title="s.Title">
                                <template class="ant-card-actions" slot="actions">
                                    <a-icon type="edit" @click="showHandleReverseHost(true, s)"/>
                                    <a-icon  type="delete" @click="deleteReverseHost(s.Id)"/>
                                </template>
                                <p style="word-wrap: break-word"><strong>是否开启反向代理</strong>: {{s.EnableReserve}}</p>
                                <div><strong>图片缩率参数是否在路径内</strong>: {{s.ReserveRule.ImgThumbInPath}}</div>
                                <div><strong>图片后台管理展示参数规则</strong>: {{s.ReserveRule.ImgShowRule}}</div>
                                <div><strong>图片发布时下载展示参数规则</strong>: {{s.ReserveRule.ImgDownloadRule}}</div>
                                <div><strong>视频展示参数规则</strong>: {{s.ReserveRule.VideoRule}}</div>
                                <div><strong>反向代理Header</strong>: {{s.HeaderStr}}</div>
                            </a-card>
                        </a-spin>
                    </a-col>
                </a-row>
            </div>
        </content-item>
        <a-modal title="添加host" :visible="addReverseHostVisible" :footer="null"
                 @cancel="()=> this.addReverseHostVisible = false">
            <a-form>
                <a-form-item label="host">
                    <a-input v-model="ReverseHostForm.Title" type="string">
                    </a-input>
                </a-form-item>
                <a-form-item label="是否开启反向代理">
                    <a-switch v-model="ReverseHostForm.EnableReserve"></a-switch>
                </a-form-item>
                <a-form-item label="图片缩率参数是否在路径内">
                    <a-switch v-model="ReverseHostForm.ReserveRule.ImgThumbInPath"></a-switch>
                </a-form-item>

                <a-form-item label="图片后台管理展示参数规则">
                    <a-input v-model="ReverseHostForm.ReserveRule.ImgShowRule">
                    </a-input>
                </a-form-item>

                <a-form-item label="图片发布时下载展示参数规则">
                    <a-input v-model="ReverseHostForm.ReserveRule.ImgDownloadRule">
                    </a-input>
                </a-form-item>
                <a-form-item label="视频展示参数规则">
                    <a-input v-model="ReverseHostForm.ReserveRule.VideoRule">
                    </a-input>
                </a-form-item>

                <a-form-item label="反向代理Header(json格式)">
                    <a-textarea v-model="ReverseHostForm.HeaderStr" placeholder="script content"
                                :autosize="{ minRows: 6, maxRows:10 }"></a-textarea>
                </a-form-item>

                <a-form-item>
                    <a-button type="primary" html-type="submit" class="login-form-button"
                              :loading="addReverseHostLoading"
                              @click.prevent="handleReverseHost">提交
                    </a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import {Button, Icon, Modal, Form, Input, List, Card, Row, Col, Switch,Spin} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {
        listReverseConfigApi,
        addReverseConfigApi,
        deleteReverseConfigApi,
        updateReverseConfigApi
    } from "../../api/reverse"

    export default {
        name: "ConfigReverseHost",
        components: {
            ContentItem,
            [Switch.name]: Switch,
            [Spin.name]: Spin,
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
            this.listReverseHost()
        },
        data() {
            return {
                addReverseHostVisible: false,
                addReverseHostLoading: false,
                ReverseHost: [],
                defaultReverseHostForm: {
                    Id: 0,
                    Title: "www.qiniu.com",
                    EnableReserve: true,
                    Header: {
                        "Referer": "https://www.qiniu.com",
                        "Sec-Fetch-Mode": "no-Crawler",
                        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",
                    },
                    HeaderStr: '{"Referer": "https://www.qiniu.com", "Sec-Fetch-Mode": "no-Crawler", "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",}',
                    ReserveRule: {
                        ImgThumbInPath: false,
                        ImgDownloadRule: "imageMogr2/thumbnail/!88p",// scheme host path param
                        ImgShowRule: "imageView2/2/w/276",
                        VideoRule: "",
                    },
                },
                ReverseHostForm: {
                    Id: 0,
                    Title: "",
                    EnableReserve: true,
                    Header: {},
                    HeaderStr: "",
                    ReserveRule: {
                        ImgThumbInPath: false,
                        ImgDownloadRule: "imageMogr2/thumbnail/!88p",// scheme host path param
                        ImgShowRule: "imageView2/2/w/276",
                        VideoRule: "",
                    },
                }
            }
        },
        methods: {
            showHandleReverseHost(isUpdate, data) {
                this.addReverseHostVisible = true;
                if (!isUpdate) {
                    this.ReverseHostForm = Object.assign({}, this.defaultReverseHostForm);
                    return
                }
                this.ReverseHostForm = data
            },
            handleReverseHost() {
                if (this.addReverseHostLoading) {
                    return
                }
                this.addReverseHostLoading = true;
                if (this.ReverseHostForm.Id === 0) {
                    this.addReverseHost()
                } else {
                    this.updateReverseHost()
                }
            },
            listReverseHost() {
                listReverseConfigApi().then(res => {
                    if (!res) {
                        return
                    }
                    const result = res.data;
                    const list = [];
                    for (let i in result) {
                        let s = JSON.parse(result[i].Value);
                        s.Id = result[i].Id;
                        s.Loading = false;
                        list.push(s)
                    }
                    this.ReverseHost = list
                })
            },
            addReverseHost() {
                addReverseConfigApi(this.createFormData()).then(res => {
                    this.addReverseHostLoading = false;
                    if (!res) {
                        return
                    }
                    this.addReverseHostVisible = false;
                    this.listReverseHost()
                })
            },
            updateReverseHost() {
                updateReverseConfigApi(this.createFormData()).then(res => {
                    this.addReverseHostLoading = false;
                    if (!res) {
                        return
                    }
                    this.addReverseHostVisible = false;
                    this.listReverseHost()
                })
            },
            deleteReverseHost(id) {
                const index = this.ReverseHost.findIndex((c) => {
                    return c.Id === id;
                });
                if (this.ReverseHost[index].Loading){
                    return
                }
                this.ReverseHost[index].Loading = true;

                const form = new FormData;
                form.append("id", id);
                deleteReverseConfigApi(form).then(res => {
                    this.ReverseHost[index].loading = false;
                    if (!res) {
                        return
                    }
                    this.ReverseHost.splice(index, 1);
                })
            },
            createFormData() {
                const form = new FormData();
                form.append("id", this.ReverseHostForm.Id);
                form.append("title", this.ReverseHostForm.Title);
                form.append("value", JSON.stringify(this.ReverseHostForm));
                return form
            },
        }
    }
</script>

<style scoped>
    .ReverseHost-card {
        margin-bottom: 16px;
    }
</style>
