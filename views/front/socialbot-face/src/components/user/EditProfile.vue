<template>
    <div class="model-box">
        <div class="form-group">
            <label class="upload-box">
                <span class="upload">
                    <van-image v-if="showAvatar" class="relative-img" width="128" height="128" fit="cover" :src="showAvatar" ></van-image>
                    <img v-else class="relative-img" src="../../assets/img/avatar.gif" alt="avatar">
                </span>
                <van-uploader class="upload-btn" :before-read="beforeRead"  :after-read="afterRead" :max-count="1" :preview-image="false">
                    <img class="relative-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAMAAABg3Am1AAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAByUExURUdwTACIzQCHzgCW2QCHzACGzQD//wCHzQCOzQCHzQCN0wCGzQCGzQCJ1QCHzQCGzQCHzQCM0gCHzQCR0gCHzQCHzQCMzgCq5ACHzQCHzwCHzQCHzQCLzACHzQCHzQCHzQCHzACIzgCHzACJzQCIzQCGzGRjDXMAAAAldFJOUwBldwjI8wF/Gb0O3foS19DMHMEW/sQfBItLuoAlrHRd6jv3MJ7YRSqoAAABeUlEQVRIx7XWx5qDMAxFYYVmY2pCei9z3v8VZwGeAAFiFqO1/nz4KiCLSLQIQhwqDBaRiIiOca5Yi0Qz+iGOZAGpp8WhtJfCQgLwxLE8CCQE7Qo0hAKISJ2VzWG0AAtsVrF2A++s4sgJ1Fk1ObgAm5UHgROwWWkI/wfMfqTZh54d6+TgttsBMPHXWCbJcgCM1jKBRLsDnQPklSuocgCoxXdQ5YBSQJG5gKwAlIgCyuw7sP2NMN+AKW1/LQLzCc7b4f5abD7A6/43IzFBu78WffAIwQqz6fZb0QaHJ9ip+qt+fyNa4JbCGsi1iB9/9tfiDa5reB5OQF75x6H+7hz2wP1R/0o+0t8Cux8gOYuInABATb5A/gUodOts+8k3zqyA2LTPVmQToCqAy66bRpmNgwTYR/38SjMKYH0dSDwwoyC9iTiJBoSHgcdVwMYMg9dghApY+eOTdhLA1FJUQOz3l+LU2lXA0e+t3cnF3hHNt3321WH25WTu9ecXNRJu8OBc1x4AAAAASUVORK5CYII=" alt="avatar">
                </van-uploader>
            </label>
        </div>
        <div class="form-group">
            <label>
                <input class="nickname" v-model="user.nickname" type="text"  placeholder="nickname" />
            </label>
        </div>
        <div class="form-group">
            <label>
                <textarea v-model="user.intro" rows="3" placeholder="intro" class="intro"></textarea>
            </label>
        </div>
        <div class="form-group submit">
            <van-button  :loading="loading"  loading-type="spinner" loading-text="submit..." type="info" @click="submit" block round >Submit</van-button>
        </div>
        <van-icon class="cross" name="cross" @click="$emit('modelAction', 'close')" ></van-icon>
    </div>
</template>

<script>
    import { Uploader,Notify, Button,Icon,Image } from 'vant';
    import {editProfile} from "../../api/user";
    import { mapActions } from 'vuex';

    export default {
        name: "EditProfile",
        components: {
            [Uploader.name]: Uploader,
            [Notify.name]: Notify,
            [Button.name]: Button,
            [Icon.name]: Icon,
            [Image.name]:Image
        },
        props:{
            user: {
                type: Object,
                default: () => {
                    return {
                        avatar:"",
                        intro: "",
                        nickname:""
                    }
                }
            },
        },
        data() {
            return {
                loading:false,
                showAvatar: "",
                editAvatar: {},
            }
        },
        created(){
          this.showAvatar = this.user.avatar;
        },
        methods: {
            ...mapActions({
                saveProfile : "SAVE_PROFILE"
            }),
            beforeRead(file) {
                if (file.type !== 'image/jpeg' && file.type !== 'image/png' && file.type !== 'image/gif') {
                    Notify('invalid img !');
                    return false;
                }
                return true;
            },
            afterRead(file) {
                this.showAvatar = file.content;
                this.editAvatar = file.file;
            },
            submit(){
                if (this.loading) {
                    return
                }
                this.loading = true;

                editProfile(this.user.nickname, this.user.intro, this.editAvatar).then((res) => {
                    this.loading = false;
                    if (!res){
                        return false
                    }

                    this.saveProfile(res.data);
                    this.$emit('modelAction', 'close');
                    return true
                });
            }
        }
    }
</script>

<style scoped>
    .upload-box {
        position: relative;
        display: inline-block;
    }
    .upload {
        border-radius: 50%;
        overflow: hidden;
        width: 8rem;
        display: inline-block;
        height: 8rem;
    }
    .upload-btn {
        padding: 0.6rem;
        background: #fff;
        border-radius: 50%;

        box-shadow: 0 3px 10px 0 rgba(0,0,0,.25);
        cursor: pointer;
        height: 2.4rem;
        width: 2.4rem;
        position: absolute;
        right: 0;
        bottom: 0;
        overflow: hidden;
    }


    .form-group .intro, .form-group .nickname {
        -webkit-appearance: none;
        -moz-appearance: none;
        appearance: none;
        border-radius: 0;
        border: 1px solid #ececec;
        box-sizing: border-box;
        overflow: hidden;
        padding: 1rem;
        resize: none;
        width: 100%;
    }

    .submit {
        padding: 0 10%;
    }

    @media (max-width: 576px) {
        .upload {
            width: 5rem;
            height: 5rem;
        }
        .upload-btn {
            padding: 0.3rem;
            height: 1.8rem;
            width: 1.8rem;
        }
        .form-group .intro, .form-group .nickname {
            padding: 0.5rem;
        }
    }

</style>
