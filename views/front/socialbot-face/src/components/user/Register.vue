<template>
    <div class="model-box">
        <div class="form-group">
            <h2 class="h4 mb-0">Discover the top wedding ideas</h2>
            <p>The go-to destination for discovering and connecting with different wedding.</p>
        </div>
        <van-cell-group class="form-group">
            <van-field v-model="account" type="text" placeholder="email" left-icon="manager"></van-field>
            <van-field v-model="password" type="password" placeholder="password" left-icon="lock"></van-field>
        </van-cell-group>

        <div class="form-group">
            <van-button :loading="loginLoading"  loading-type="spinner" loading-text="register..." type="info" @click="register" block round >Register</van-button>
        </div>
        <van-icon class="cross" name="cross" @click="$emit('modelAction', 'close')" ></van-icon>
        <!--<div class="form-group text-center">
            <span class="small text-muted">Do not have an account?</span>Signup
        </div>-->
    </div>
</template>

<script>
    import {Field, Icon, CellGroup, Button} from 'vant';
    import {registerApi} from '../../api/user'

    import { mapActions } from 'vuex';

    export default {
        name: 'login',
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
            register() {
                if (this.loginLoading) {
                    return
                }
                this.loginLoading = true;
                registerApi(this.account, this.password, this.type).then((res) => {
                    this.loginLoading = false;
                    if (!res){
                        return false
                    }
                    this.loginStore(res.data);
                    this.$emit('modelAction', 'profile');
                    return true
                });
            },
            ...mapActions({
                loginStore : "LOGIN"
            }),
        }
    }
</script>

<style scoped>
</style>
