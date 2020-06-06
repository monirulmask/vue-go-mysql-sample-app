new Vue({
    el: '#app',

    data: {
        teachers: [],
        teacherName: '',
        teacherSalary: '',
        current: -1,
        isEntered: false
    },

   
    computed: {
        computedTeachers() {
          return this.teachers
        },
        validate() {
            var isEnteredteacherName = 0 < this.teacherName.length
            this.isEntered = isEnteredteacherName
            return isEnteredteacherName
        }
    },

    created: function() {
        this.doFetchAllTeachers()
    },

    methods: {
        doFetchAllTeachers() {
            axios.get('/fetchAllTeachers')
            .then(response => {
                if (response.status != 200) {
                    throw new Error('Response Now Found!!!')
                } else {
                    var resultTeachers = response.data
                    this.teachers = resultTeachers
                }
            })
        },
        doFetchTeacher(teacher) {
            axios.get('/fetchTeacher', {
                params: {
                    teacherID: teacher.id
                }
            })
            .then(response => {
                if (response.status != 200) {
                    throw new Error('Response Now Found!!!')
                } else {
                    var resultTeacher = response.data
                    var index = this.teachers.indexOf(teacher)
                    this.teachers.splice(index, 1, resultTeacher[0])
                }
            })
        },
        doAddTeacher() {
            const params = new URLSearchParams();
            params.append('teacherName', this.teacherName)
            params.append('teacherSalary', this.teacherSalary)

            axios.post('/addTeacher', params)
            .then(response => {
                if (response.status != 200) {
                    throw new Error('Response Now Found!!!')
                } else {
                    this.doFetchAllTeachers()
                    this.initInputValue()
                }
            })
        },
        doChangeTeacherState(teacher) {
            const params = new URLSearchParams();
            params.append('teacherID', teacher.id)
            params.append('teacherState', teacher.state)

            axios.post('/changeStateTeacher', params)
            .then(response => {
                if (response.status != 200) {
                    throw new Error('Response Now Found!!!')
                } else {
                    this.doFetchTeacher(teacher)
                }
            })
        },
        doDeleteTeacher(teacher) {
            const params = new URLSearchParams();
            params.append('teacherID', teacher.id)

            axios.post('/deleteTeacher', params)
            .then(response => {
                if (response.status != 200) {
                    throw new Error('Response Now Found!!!')
                } else {
                    this.doFetchAllTeachers()
                }
            })
        },
        initInputValue() {
            this.current = -1
            this.teacherName = ''
            this.teacherSalary = ''
        }
    }
})
