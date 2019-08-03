<template>
    <div class="login-form">
        <a-form id="components-form-demo-normal-login" :form="form" @submit="handleSubmit">
            <a-form-item>
                <a-input v-decorator="['account',{ rules: [{ required: true, message: '请输入手机号!' }] }]" type="integer" placeholder="手机号">
                    <a-icon slot="prefix" type="user" style="color: rgba(0,0,0,.25)"/>
                </a-input>
            </a-form-item>
            <a-form-item>
                <a-input v-decorator="['password',{ rules: [{ required: true, message: '请输入密码!' }] }]" type="password" placeholder="密码">
                    <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)"/>
                </a-input>
            </a-form-item>
            <a-form-item>
                <a-button type="primary" html-type="submit" class="login-form-button" :loading="submitLoading">登录</a-button>
            </a-form-item>
        </a-form>
    </div>
</template>

<script>
    import { Form, Button,Input, Icon} from 'ant-design-vue'
    import {loginApi} from "../api/user";
    import {mapMutations} from "vuex"

    export default {
        name: 'Login',
        data () {
            return {
                submitLoading: false
            }
        },
        beforeCreate () {
            this.form = this.$form.createForm(this)
        },
        components:{
            [Button.name]: Button,
            [Input.name]: Input,
            [Form.name]: Form,
            [Form.Item.name]: Form.Item,
            [Icon.name]: Icon
        },
        methods: {
            ...mapMutations({
                loginSaveUser : "loginSaveUser"
            }),
            handleSubmit (e) {
                e.preventDefault()
                this.form.validateFields((err, values) => {
                    if (err) {
                        return false
                    }
                    this.submitLoading = true;
                    loginApi(values.account, values.password, 'email').then(response => {
                        console.log("----->",response)
                        this.submitLoading = false;
                        if (!response){
                            return
                        }
                       this.loginSaveUser(response.data);
                        this.$router.push('/dashboard')
                    })

                })
            }
        }
    }
</script>
<style>
    .login-form {
        max-width: 300px;
        margin:3rem  auto ;
    }
    .login-form-forgot {
        float: right;
    }
    .login-form-button {
        width: 100%;
    }
</style>
