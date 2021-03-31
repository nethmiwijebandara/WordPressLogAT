import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-detail-dialog',
  templateUrl: './detail-dialog.component.html',
  styleUrls: ['./detail-dialog.component.sass']
})
export class DetailDialogComponent implements OnInit {

  constructor(
    @Inject(MAT_DIALOG_DATA) public data,

  ) { }

  ngOnInit(): void {
  }

}
