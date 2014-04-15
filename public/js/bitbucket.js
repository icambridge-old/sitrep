window.RepoInfo   = Backbone.Model.extend({urlRoot : '/ajax/bitbucket/pullrequests'})

window.JenkinsJobsListView = Backbone.View.extend({

    tagName:'div',

    initialize:function () {
        this.collection.bind("reset", this.render, this);
    },

    render:function (eventName) {
        $("#loading").hide()
        this.collection.each(function (job) {
                $(this.el).append(new JenkinsJobsItemView({model:job}).render().el);
            }, this);
        return this;
    }
});


window.JenkinsJobsItemView = Backbone.View.extend({
    tagName:"div",

    template:_.template($('#jenkins-job').html()),

    render:function (eventName) {
        $(this.el).html(this.template(this.model.toJSON()));
        
        return this;
    }
});

window.RepoInfoView = Backbone.View.extend({

    tagName:"div",

    template:_.template($('#repo-info').html()),

    initialize:function () {

         this.listenTo(this.model,'change', this.render);
    },

    render:function (eventName) {
        json = this.model.toJSON()
        $('.to-be-used').html(this.template(json));
        
        return this;
    }
});
var AppRouter = Backbone.Router.extend({
	routes: {
		""			: "list",
		"job/:name" : "jobDetails",
		"merge/:name/:id" : "merge",
		"build/:name/:id" : "build"
	},

    list:function() {

        console.log("Home");
    },

    initialize:function () {

        this.jenkinsJobsView = new JenkinsJobsListView({collection:jobs});
        this.jenkinsJobsView.render()
        $('#job-list').html(this.jenkinsJobsView.el);


    },

    merge: function (name, id) {

        url = "/ajax/bitbucket/pullrequests/" + name + "/merge/" + id
        $.get( url, function( data ) {
            if (data.status == "success") {
                $("tr#" + name + "-" + id).remove()
            } else {
                alert("Something went wrong! Create a github issue!")
            }
        },"json");
        this.navigate("/job/" + name, {trigger: false})
    },
    build: function (name, branch) {

        $("td#" + name + "-" + branch).html("Building")
        url = "/jenkins/build/" + name + "/" + branch
        $.get( url, function( data ) {
            // Do stuff here
        },"json");

        this.navigate("/job/" + name, {trigger: false})
    },

    jobDetails:function (nameInput) {
        $("section.list-group-item").addClass("hidden");
        $("section.list-group-item").removeClass("to-be-used");
        $("a.list-group-item").removeClass("active");
        repoName = nameInput.toLowerCase()
        repo = new RepoInfo({id: repoName})
        repoView = new RepoInfoView({model: repo})
        repo.fetch()
        $("#job-" + nameInput).addClass("to-be-used");
        $("#job-" + nameInput).html('<div class="progress progress-striped active"><div class="progress-bar"  role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100" style="width: 100%"><span class="sr-only">100% Complete</span></div></div>');
        $("#job-" + nameInput).removeClass("hidden");
        $("#job-" + nameInput).show();
        $("#job-" + nameInput + "-link").addClass("active")
    },

});

var app = new AppRouter();
Backbone.history.start();
