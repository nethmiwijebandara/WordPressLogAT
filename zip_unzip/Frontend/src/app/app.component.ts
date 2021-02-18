import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import {FileuploadService} from '../app/services/fileupload.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit{
  mediaSub:Subscription;
  fileToUpload: File = null; 
  loading: Boolean = false; 
  successMsg = ''
  title: any
  selectedFile= null;
  name=null;
  constructor(public fileUpload : FileuploadService){}
  ngOnInit(){
    
  }

  handleFileInput(files: FileList) {
 
    this.fileToUpload = files.item(0);
    console.log(this.fileToUpload, "select file")
}

  uploadFileToActivity() {
    this.loading = true
    console.log(this.fileToUpload, "submit button")
    this.fileUpload.fileUpload(this.fileToUpload).subscribe(res=>{
      this.successMsg = res['result']
      this.loading = false
      window.alert( this.successMsg)
    
    })

  }
}


  
  

