<template>
    <div class="config">
        <a-row :gutter="50">
            <a-col :sm="1" :md="12">
                <a-spin :spinning="!website">
                    <a-card title="网站设置" class="card">
                        <a-form-item  label="网站名称">
                            <a-input v-model="website.HostName" type="string"></a-input>
                        </a-form-item>
                        <a-form-item>
                            <a-button type="primary" html-type="submit" class="login-form-button" :loading="upWebsiteLoading" @click.prevent="upWebsite">更新</a-button>
                        </a-form-item>
                    </a-card>
                </a-spin>
            </a-col>
            <a-col :sm="1" :md="12">
                <a-spin :spinning="!cors">
                    <a-card title="跨域设置" class="card">
                        <a-form-item  label="是否开启">
                            <a-switch v-model="cors.Enable" ></a-switch>
                        </a-form-item>
                        <a-form-item  label="AllowCredentials">
                            <a-switch v-model="cors.AllowCredentials" ></a-switch>
                        </a-form-item>
                        <a-form-item  label="MaxAge">
                            <a-input v-model="cors.MaxAge" type="number"></a-input>
                        </a-form-item>
                        <a-form-item  label="allowOrigins (json格式)">
                            <a-textarea v-model="allowOrigins" type="string"></a-textarea>
                        </a-form-item>
                        <a-form-item  label="allowOrigins (json格式)">
                            <a-textarea v-model="allowMethods" type="string"></a-textarea>
                        </a-form-item>
                        <a-form-item  label="allowOrigins (json格式)">
                            <a-textarea v-model="allowHeaders" type="string"></a-textarea>
                        </a-form-item>
                        <a-form-item>
                            <a-button type="primary" html-type="submit" class="login-form-button" :loading="upCorsLoading" @click.prevent="upCors">更新</a-button>
                        </a-form-item>
                    </a-card>
                </a-spin>
            </a-col>
        </a-row>
    </div>
</template>

<script>
    import {Button, Icon, Form,Input, Row,Col,Card, Switch,Spin,message} from 'ant-design-vue'
    import {BaseConfigApi, updateConfigApi} from "../../api/config"
    export default {
        name: "Config",
        components: {
            [Spin.name]: Spin,
            [Card.name]: Card,
            [Button.name]: Button,
            [Icon.name]: Icon,
            [Form.name]: Form,
            [Form.Item.name]: Form.Item,
            [Input.name]: Input,
            [Input.TextArea.name]: Input.TextArea,
            [Switch.name]: Switch,
            [Row.name]: Row,
            [Col.name]: Col,
        },
        created(){
            this.baseConfig()
        },
        data() {
            return {
                upWebsiteLoading: false,
                upCorsLoading: false,
                website:false,
                storage:false,
                cors:false,
                reserveHostRules:false,
                allowOrigins:"[]",
                allowMethods:"[]",
                allowHeaders:"[]",
            }
        },
        methods:{
            baseConfig(){
                BaseConfigApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.website = res.data.website;
                    this.storage = res.data.storage;
                    this.cors = res.data.cors;
                    this.allowOrigins = JSON.stringify(this.cors.AllowOrigins);
                    this.allowMethods = JSON.stringify(this.cors.AllowMethods);
                    this.allowHeaders = JSON.stringify(this.cors.AllowHeaders);
                })
            },
            upWebsite(){
                if (this.upWebsiteLoading){
                    return
                }
                this.upWebsiteLoading = true;
                const form = new FormData();
                form.append("key", "website");
                form.append("title", "website");
                form.append("value", JSON.stringify(this.website));
                updateConfigApi(form).then(res => {
                    this.upWebsiteLoading = false;
                    if (res) {
                        message.success('更新成功');
                    }
                })
            },
            upCors(){
                if (this.upCorsLoading){
                    return
                }
                this.upCorsLoading = true;
                this.cors.MaxAge = parseInt(this.cors.MaxAge );
                this.cors.AllowOrigins = JSON.parse(this.allowOrigins);
                this.cors.AllowMethods = JSON.parse(this.allowMethods);
                this.cors.AllowHeaders = JSON.parse(this.allowHeaders);
                const form = new FormData();
                form.append("key", "cors");
                form.append("title", "cors");
                form.append("value", JSON.stringify(this.cors));
                updateConfigApi(form).then(res => {
                    this.upCorsLoading = false;
                    if (res) {
                        message.success('更新成功');
                    }
                })
            },
        }
    }
</script>

<style scoped>
    .config {
        padding: 10px;
    }
    .card {margin-bottom: 10px;}
</style>
