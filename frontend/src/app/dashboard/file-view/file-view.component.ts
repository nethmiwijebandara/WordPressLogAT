import { Component, OnInit } from '@angular/core';
import { PageEvent } from '@angular/material/paginator';
import { FileViewService } from '../../services/file-view.service';

@Component({
  selector: 'app-file-view',
  templateUrl: './file-view.component.html',
  styleUrls: ['./file-view.component.scss']
})
export class FileViewComponent implements OnInit {

  itemList = [];
  pagedList = [];

  constructor(private fileViewService: FileViewService) { }

  ngOnInit(): void {
    this.getFileview();
  }

  getFileview() {
    this.fileViewService.getFileView().subscribe( res => {
      this.itemList = res.body.results;
      this.pagedList = this.itemList.slice(0,9);
      console.log(this.itemList);
    }, err => {
      console.log(err);
    }
    );
  }

  onPageChange(event: PageEvent) {
    const matTable = document.getElementById('content-wrapper');
    matTable.scrollIntoView();

    let startIndex = event.pageIndex * event.pageSize;
    let endIndex = startIndex + event.pageSize;
    if (endIndex > this.itemList.length) {
      endIndex = this.itemList.length;
    }
    this.pagedList = this.itemList.slice(startIndex, endIndex);
  }
}
