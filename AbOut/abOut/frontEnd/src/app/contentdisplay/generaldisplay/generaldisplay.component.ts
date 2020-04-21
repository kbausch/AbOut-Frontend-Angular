import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-generaldisplay',
  templateUrl: './generaldisplay.component.html',
  styleUrls: ['./generaldisplay.component.scss']
})
export class GeneraldisplayComponent implements OnInit {

  // The booleansare used below to toggle to correct display.
  private outcomes = false;
  private programs = false;

  constructor(private route: ActivatedRoute) { }

  ngOnInit() {

    // The route query below will be used to toggle the correct display.
    this.route.queryParams.subscribe(params => {
      if (params.type === 'outcomes') {
        this.outcomes = true;
        this.programs = false;
      } else if (params.type === 'programs') {
        this.programs = true;
        this.outcomes = false;
      }
    });
  }

}
