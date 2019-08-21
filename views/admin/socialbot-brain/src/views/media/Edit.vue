<template>
    <div class="edit">
        <a-row :gutter="50">
            <a-spin :spinning="loadingDetail">
                <a-col :sm="1" :md="12">
                    <div class="group recommend">
                        <label class="mr8">是否推荐</label>
                        <a-switch v-model="activeRecommend"></a-switch>
                    </div>
                    <div class="group cate">
                        <a-radio-group v-model="activeCategory">
                            <a-radio v-for="c in categoryList" :key="c.Id" :value="c.Id">{{c.ShortName}}</a-radio>
                        </a-radio-group>
                    </div>
                    <div class="group tag">
                        <span :class="tagClass(tag.Id)" v-for="tag in tagList() " :key="tag.Id" @click="toggleTag(tag)">
                            {{tag.ShortName}}
                        </span>
                    </div>
                    <div class="group">
                        <a-textarea v-model="activeTitle"  placeholder="Basic usage"></a-textarea>
                    </div>
                    <div class="group">
                        <a-button :loading="editLoading" type="primary" @click="editMedia">编辑</a-button>
                    </div>
                </a-col>
                <a-col :sm="1" :md="12">
                    <img class="cover" :src="activeCover" />
                </a-col>
            </a-spin>
        </a-row>
    </div>
</template>

<script>
    import {Button, Radio, Input, Switch, Select, Row, Col, Spin,message} from 'ant-design-vue'
    import {editMediaApi, mediaDetailApi} from "../../api/media";
    import {listCategoryWithTagsApi} from "../../api/category"

    export default {
        name: "EditMedia",
        props: {
            uri: {
                Type: String,
                default: () => {
                    return ""
                }
            }
        },
        components: {
            [Spin.name]: Spin,
            [Row.name]: Row,
            [Col.name]: Col,
            [Button.name]: Button,
            [Input.name]: Input,
            [Input.TextArea.name]: Input.TextArea,
            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
            [Radio.name]: Radio,
            [Switch.name]: Switch,
            [Radio.Group.name]: Radio.Group,
        },
        data() {
            return {
                categoryList: [],
                loadingDetail: false,
                editLoading: false,
                activeTitle: "",
                activeCover:"",
                activeCategory: 0,
                activeRecommend: false,
                activeTags: [],
            }
        },
        created() {
            this.listCategory();
            this.mediaDetail()
        },
        methods: {
            toggleTag(tag) {
                const index = this.activeTags.findIndex((t) => {
                    return t === tag.Id;
                });
                if (index >= 0) {
                    this.activeTags.splice(index, 1);
                } else {
                    this.activeTags.push(tag.Id);
                }
            },
            tagClass(tagId) {
                let isSelected = this.activeTags.includes(tagId);
                return {'active': isSelected}
            },
            editMedia() {
                if (this.editLoading) {
                    return
                }
                this.editLoading = true;
                const f = new FormData;
                f.append("uri", this.uri);
                f.append("title", this.activeTitle);
                f.append("tags", this.activeTags);
                f.append("cid", this.activeCategory);
                f.append("recommend", this.activeRecommend);
                editMediaApi(f).then((res) => {
                    this.editLoading = false;
                    if (res) {
                        message.success('更新成功');
                    }
                })
            },
            tagList() {
                if (this.categoryList.length === 0 || this.activeCategory === 0) {
                    return
                }

                const index = this.categoryList.findIndex((cate) => {
                    return cate.Id === this.activeCategory;
                });

                let nowCate = this.categoryList[index];
                return nowCate.Tags
            },
            recommend(e) {
                this.activeRecommend = e.target.checked
            },
            listCategory() {
                listCategoryWithTagsApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.categoryList = res.data;
                });
            },
            mediaDetail() {
                this.loadingDetail = true;
                mediaDetailApi({uri: this.uri}).then((res) => {
                    this.loadingDetail = false;
                    if (!res) {
                        return
                    }
                    const result = res.data;
                    this.activeTitle = result.Title;
                    this.activeCover = result.Cover;
                    this.activeCategory = result.Cid;
                    this.activeRecommend = result.Recommend !== 0;
                    if (result.Tags.length > 0) {
                        for (let i in result.Tags) {
                            this.activeTags.push(result.Tags[i].Id)
                        }
                    }
                })
            },
        }
    }
</script>

<style scoped>
    .edit {
        padding: 15px;
        min-height: 500px;
    }
    .group {
        margin-bottom: 12px;
    }
    .cover{
        width: 100%;
    }
    .tag {
        margin-top: 8px;
    }
    .tag span {
        display: inline-block;
        cursor: pointer;
        margin-right: 0.3rem;
        background: #f0f0f0;
        border-radius: 0.8rem;
        padding: 0 0.5rem;
        font-size: 0.8rem;
        margin-bottom: 0.2rem;
    }

    .tag span.active {
        color: #ff6a1e;
    }
</style>
