<template>
    <div class="action-nav" ref="actionNav">
        <avatar @click.native="toggleNav" :urlStr="user.avatar"  ></avatar>
        <div v-show="show" class="nav-item">
            <van-cell @click="$emit('modelAction','profile')" title="Edit Profile"></van-cell>
            <van-cell @click="logout" title="Logout"></van-cell>
        </div>
        <div class="cover" @click="closeNav"></div>
    </div>
</template>

<script>
    import Avatar from './Avatar'
    import { Cell } from 'vant';

    export default {
        name: "ActionNav",
        components: {
            Avatar,
            [Cell.name]: Cell
        },
        props:{
            user: {
                type: Object,
                default: () => {
                    return {
                        avatar:"",
                    }
                }
            },
        },
        data(){
          return {
              show: false
          }
        },
        mounted(){
            this.listen(window, 'click', (event) => {
                if(!this.$refs.actionNav.contains(event.target)){
                    this.show = false;
                }
                // eslint-disable-next-line no-console
                /*console.log(1111111)*/
            })
        },
        beforeDestroy() {
            if (this._eventRemovers) {
                this._eventRemovers.forEach((eventRemover) => {
                    eventRemover.remove()
                })
            }
        },
        methods:{
            listen(target, eventType, callback) {
                if (!this._eventRemovers) {
                    this._eventRemovers = []
                }
                target.addEventListener(eventType, callback);
                this._eventRemovers.push({
                    remove: function() {
                        target.removeEventListener(eventType, callback)
                    }
                })
            },
            toggleNav(){
               this.show = !this.show
            },
            closeNav(){
                this.show= false
            },
            logout(){
                this.$store.dispatch('LOGOUT')
            }
        }
    }
</script>

<style scoped>
    .action-nav {
        position: relative;
    }
    .action-nav .nav-item {
        position: absolute;
        top: 3.2rem;
        right: 0;
        z-index: 1003;
        width: 14rem;
        -webkit-filter: drop-shadow(0 2px 4px hsla(0,0%,75.3%,.5));
        filter: drop-shadow(0 2px 4px rgba(192,192,192,.5));
    }
    .cover {display: block}
</style>
