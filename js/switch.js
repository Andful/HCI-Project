function Switcher(dom_obj) {
	//this.parent = $('.switch-content')
	this.element = $(dom_obj);
	this.element.children().hide();
	this.element.children().eq(0).show();
	//this.elements.hide()
	//this.elements[2].show()

	this.switch = function(elem){
		this.element.children().hide(500);
		this.element.children(elem).show(500);
	}
	
}

switcherSetting = new Switcher("#switch-content-setting");
switcherSchedule = new Switcher("#switch-content-schedule");