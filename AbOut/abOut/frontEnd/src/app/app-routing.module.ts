import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { GeneraldisplayComponent } from './contentdisplay/generaldisplay/generaldisplay.component';
import { ViewdisplayComponent } from './contentdisplay/viewdisplay/viewdisplay.component';
import { ViewprogramdisplayComponent } from './contentdisplay/viewprogramdisplay/viewprogramdisplay.component';
import { OutcomesSidenavbarComponent } from './sidenavbar/outcomes-sidenavbar/outcomes-sidenavbar.component';
import { ProgramsSidenavbarComponent } from './sidenavbar/programs-sidenavbar/programs-sidenavbar.component';
import { CreateoutcomedisplayComponent } from './contentdisplay/createoutcomedisplay/createoutcomedisplay.component';

const routes: Routes = [
    {path: 'programs-sidebar', component : ProgramsSidenavbarComponent, outlet: 'sidebar'},
    {path: 'outcomes-sidebar', component : OutcomesSidenavbarComponent, outlet: 'sidebar'},
    {path: 'generaldisplay', component: GeneraldisplayComponent, runGuardsAndResolvers: 'always'},
    {path: 'viewdisplay', component: ViewdisplayComponent, runGuardsAndResolvers: 'always'},
    {path: 'createoutcomedisplay', component: CreateoutcomedisplayComponent},
    {path: 'viewprogramdisplay', component: ViewprogramdisplayComponent, runGuardsAndResolvers: 'always'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {onSameUrlNavigation: 'reload'})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
