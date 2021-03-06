// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package trello

const boardsTemplate = `
<div class="section-trello-render">
	{{if gt (len .Boards) 0}}
		<div class="heading">Boards</div>
		<p>There are {{len .Boards}} boards, {{.ListTotal}} lists, {{.CardTotal}} cards and {{len .MemberBoardAssign}} members. Activity since {{.Since}}</p>
		<div class="section-trello-render">
			<table class="trello-table" class="width-100">
				<tbody class="trello">
				{{range $b := .Boards}}
					<tr>
						<td>
							<a href="{{ $b.Board.URL }}">
								<div class="trello-board" style="background-color: {{$b.Board.Prefs.BackgroundColor}}">
									{{$b.Board.Name}}
									<span>{{$b.Board.OrgName}}</span>
								</div>
							</a>
						</td>
						<td>
							<div class="board-summary">
								<!-- {{ len $b.Actions }}{{if eq 1 (len $b.Actions)}} action {{else}} actions {{end}} -->
							</div>
							<span class="board-meta">
								{{range $idx, $act := $b.ActionSummary}}{{if ne $idx 0}}{{- ","}} {{end}}{{$act.Count}} {{$act.Name -}}{{if ne 1 $act.Count}}{{"s" -}}{{end}}{{end}}{{if gt (len $b.Archived) 0}}, {{end}}
								{{if gt (len $b.Archived) 0}}
								 	{{len $b.Archived}} {{if eq 1 (len $b.Archived)}}card {{else}} cards {{end}}archived
								{{else}}
									{{if eq (len $b.ActionSummary) 0}}
										no activity
									{{end}}
								{{end}}
								<br>
							</span>
						</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	{{end}}
</div>
`
