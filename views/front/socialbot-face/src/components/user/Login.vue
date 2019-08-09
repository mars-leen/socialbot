<template>
    <div class="model-box">
        <div class="form-group">
            <h2 class="h4 mb-0">Welcome Back!</h2>
            <p>Login to manage your account.</p>
        </div>
        <van-cell-group class="form-group">
            <van-field v-model="account" type="text" placeholder="email" left-icon="manager"></van-field>
            <van-field v-model="password" type="password" placeholder="password" left-icon="lock"></van-field>
        </van-cell-group>

        <div class="form-group">
            <van-button :loading="loginLoading"  loading-type="spinner" loading-text="login..." type="info" @click="login" block round >Login</van-button>
        </div>
        <van-icon class="cross" name="cross" @click="$emit('modelAction', 'close')" ></van-icon>
        <!--<div class="form-group text-center">
            <span class="small text-muted">Do not have an account?</span>Signup
        </div>-->
    </div>
</template>

<script>
    import {Field, Icon, CellGroup, Button} from 'vant';
    import {loginApi} from '../../api/user'

    import { mapActions } from 'vuex';

    export default {
        name: 'Login',
        data() {
            return {
                loginLoading: false,
                account: "",
                password: "",
                type: "email",
            }
        },
        components: {
            [Field.name]: Field,
            [Icon.name]: Icon,
            [CellGroup.name]: CellGroup,
            [Button.name]: Button,
        },
        methods: {
            login() {
                if (this.loginLoading) {
                    return
                }
                this.loginLoading = true;
                loginApi(this.account, this.password, this.type).then((res) => {
                    this.loginLoading = false;
                    if (!res){
                        return false
                    }
                    this.loginStore(res.data);
                    this.$emit('modelAction', 'close');
                    return true
                }).catch(err => {
                    this.loginLoading = false;
                });
            },
            ...mapActions({
                loginStore : "LOGIN",
            }),
        }
    }
</script>

<style scoped>
</style>
