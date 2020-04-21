import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Program } from '../../models/program.model';
import { ToastService } from '../../services/toast/toast.service';

@Component({
  selector: 'app-viewprogramdisplay',
  templateUrl: './viewprogramdisplay.component.html',
  styleUrls: ['./viewprogramdisplay.component.scss']
})
export class ViewprogramdisplayComponent implements OnInit {

  private program = {
    abbrev : '',
    name: '',
    current_semester: ''
  };

  constructor(private route: ActivatedRoute, private http: HttpClient, private toastService: ToastService) { }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.showProgram(params.abbrev);
    });
  }

  // getOutcome gets the JSON from the backend and stores it in the outcome object for display.
  // It also displays an error message via the toast service if one occurs.
  private showProgram(abbrev: string) {
    this.http.get<Program>('programs/' + abbrev).subscribe(
      (response) => {
        this.program.abbrev = response.abbrev;
        this.program.name = response.name;
        this.program.current_semester = response.current_semester;
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    );
  }
}
