<template>
    <div>
        <a-form class="add-product" :form="form" @submit="handleSubmit" >
            <a-row :gutter="50">
                <a-col :sm="1" :md="12">
                    <a-form-item  label="标题">
                        <a-textarea v-decorator="['title',{ rules: [{ required: true, message: '标题!' }] }]" type="integer"
                                 placeholder="标题">
                        </a-textarea>
                    </a-form-item>
                    <a-form-item  label="商品链接">
                        <a-input v-decorator="['detail_link',{ rules: [{ required: true, message: '商品链接!' }] }]"
                                 type="string" placeholder="商品链接">
                        </a-input>
                    </a-form-item>
                    <a-form-item  label="推广链接">
                        <a-input v-decorator="['promote_link',{ rules: [{ required: true, message: '推广链接!' }] }]"
                                 type="string" placeholder="推广链接">
                        </a-input>
                    </a-form-item>
                    <a-form-item  label="选择分类">
                        <a-select  v-model="activeCategory">
                            <a-select-option v-for="cate in category" :key="cate.Id" :value="cate.Id">{{cate.Title}}</a-select-option>
                        </a-select>
                    </a-form-item>
                </a-col>
                <a-col :sm="1" :md="12">
                    <a-form-item  label="选择标签">
                        <a-tag  v-for="tag in tagList()" :key="tag.Id" :color="tagClass(tag.Id) ? 'red': 'blue'"  @click="toggleTag(tag.Id)" >{{tag.ShortName}}</a-tag>
                    </a-form-item>
                    <a-form-item  label="上传媒体文件">
                        <div class="clearfix">
                            <a-upload :action="action" listType="picture-card" :fileList="fileList"
                                      @change="handleChange" @preview="handlePreview" :headers="headers">
                                <div v-if="fileList.length < 10">
                                    <a-icon type="plus"/>
                                    <div class="ant-upload-text">Upload</div>
                                </div>
                            </a-upload>
                        </div>
                    </a-form-item>
                    <a-form-item>
                        <a-button type="primary" html-type="submit" class="login-form-button" :loading="submitLoading">提交</a-button>
                    </a-form-item>
                </a-col>
            </a-row>
        </a-form>
    </div>

</template>

<script>
    import {Form, Button, Input, Icon, Row, Col, Upload, notification, Select, Tag} from 'ant-design-vue'
    import AFormItem from "ant-design-vue/es/form/FormItem";
    import api from "../../api/index"
    import {listCategoryWithTagsApi} from "../../api/category"
    import {addCommissionProductApi} from "../../api/media"

    export default {
        name: "addCommissionProduct",
        components: {
            AFormItem,
            [Upload.name]: Upload,
            [Row.name]: Row,
            [Col.name]: Col,
            [Tag.name]: Tag,
            [Button.name]: Button,
            [Input.name]: Input,
            [Input.TextArea.name]: Input.TextArea,
            [Form.name]: Form,
            [Form.Item.name]: Form.Item,
            [Icon.name]: Icon,
            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
        },
        data: () => ({
            submitLoading: false,
            previewVisible: false,
            action: api.UploadSingle,
            fileList: [],
            headers: {},
            category: [],
            activeMedias: [],
            activeTags:[],
            activeCategory: 0,
        }),
        created() {
            this.listCategory();
            this.headers = {'Authorization': this.$store.getters.Token()};
        },
        beforeCreate() {
            this.form = this.$form.createForm(this);
        },
        methods: {
            handleSubmit(e) {
                e.preventDefault();
                if (this.submitLoading) {
                    return;
                }
                this.form.validateFields((err, values) => {
                    if (err) {
                        console.log('Received values of form: ', err);
                        return
                    }
                    this.submitLoading = true;
                    const param = {...values};
                    param.tags = this.activeTags.join(',');
                    param.cid = this.activeCategory;
                    param.medias = this.activeMedias.join(',');
                    addCommissionProductApi(param).then(res=>{
                        this.submitLoading = false;
                        if (!res){
                            return
                        }
                        notification.error({
                            message: 'success',
                            description: "成功啦"
                        });
                        window.location.reload()
                    })
                });
            },

            handleChange({fileList}) {
                const ms = [];
                for (let i in fileList) {
                    if (fileList[i].status === "uploading") {
                        break;
                    }
                    if (fileList[i].response.code && fileList[i].response.code !== 0) {
                        notification.error({
                            message: 'Error',
                            description: fileList[i].response.code
                        });
                        return;
                    }
                    ms.push(fileList[i].response.data.uri)
                }
                this.fileList = fileList;
                this.activeMedias = ms;
            },
            handleCancel() {
                this.previewVisible = false
            },
            handlePreview(file) {
                console.log(111)
            },
            listCategory() {
                listCategoryWithTagsApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.category = res.data;
                    if ( this.category.length >0){
                        this.activeCategory = this.category[0].Id
                    }
                })
            },
            tagList() {
                if (this.activeCategory ===0){
                    return []
                }
                const index = this.category.findIndex((c) => {
                    return c.Id === this.activeCategory
                });
                return this.category[index].Tags;
            },
            toggleTag(tagId) {
                const index = this.activeTags.findIndex((t) => {
                    return t === tagId;
                });
                if (index >= 0){
                    this.activeTags.splice(index, 1);
                }else{
                    this.activeTags.push(tagId);
                }
            },
            tagClass(tagId) {
                return this.activeTags.includes(tagId);
            },
        },
    };
</script>
<style>
    .add-product {
        margin: 0 auto;
        padding: 20px;
    }

    .form-group {
        margin-bottom: 6px;
    }
    .tag-span{
        display: inline-block;
        cursor: pointer;
        margin-right: 0.3rem;
        background: #f0f0f0;
        border-radius: 0.8rem;
        padding: 0 0.5rem;
        font-size: 0.8rem;
        margin-bottom: 0.2rem;
    }
    .tag-span.active {
        color: #ff6a1e;
    }
</style>
