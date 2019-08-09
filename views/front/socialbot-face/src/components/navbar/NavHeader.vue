<template>
    <div class="nav-header container">
        <router-link to="/" class="logo"><strong>Century.Wedding</strong></router-link>
        <div class="nav-user">
            <div v-if="!getUser()" class="sign-item">
                <span class="login" @click="showModal('login')">login</span>
                <van-button type="info"  size="small" @click="showModal('register')">register</van-button>
            </div>
            <div v-if="getUser()">
                <action-nav :user="getUser()"  @modelAction="showModal"></action-nav>
            </div>
        </div>

        <van-dialog   class="model registerModel" cancel-button-text="Cancel" confirm-button-text="Confirm" v-model="modal.show"   :show-cancel-button="false" :show-confirm-button="false" >
            <register v-if="modal.registerShow" @modelAction="showModal"></register>
            <login  v-if="modal.loginShow" @modelAction="showModal" ></login>
            <edit-profile :user="getUser()"  v-if="getUser() && modal.editShow" @modelAction="showModal" ></edit-profile>
        </van-dialog>
    </div>
</template>

<script>
    import { Dialog, Button } from 'vant';
    import Register from "../../components/user/Register"
    import Login from "../../components/user/Login"
    import EditProfile from "../../components/user/EditProfile"
    import ActionNav from "../../components/user/ActionNav"
    import { mapGetters, mapMutations } from 'vuex';
    export default {
        name: "NavHeader",
        components:{
            Register,
            Login,
            EditProfile,
            ActionNav,
            [Dialog.Component.name]: Dialog.Component,
            [Button.name]:Button,
        },
        data(){
            return {
                modal: this.getModal()
            }
        },
        methods:{
            ...mapGetters({
                getUser : "GET_USER",
                getModal: "GET_MODAL"
            }),
            ...mapMutations({
                showModal:"SHOW_MODAL",
            }),
        },
    }
</script>

<style scoped>
    .nav-header {
        padding: 1rem 0;
    }
    .logo{
        display: inline-block;
        margin: 0 auto;
        height: 50px;
        font-size: 20px;
        line-height: 50px;
    }
    .login {
        color: #6e6e6e;
        cursor: pointer;
        font-size: 0.8rem;
        font-weight: 700;
        margin-right: 1rem;
    }
    .model {  max-width: 40rem;    width: 90%;}
    @media (max-width: 576px) {
        .nav-header {
            padding: 0.6rem 0.8rem;
        }
        .logo{
            margin: 0;
            height: 30px;
            font-size: 18px;
            line-height: 30px;
        }
        .login {
            font-size: 0.8rem;
            margin-right: 0.6rem;
        }
    }
</style>
