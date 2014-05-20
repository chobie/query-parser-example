#include <stdio.h>
#include "query_types.h"
#include "y.tab.h"

int main() {
	//query q = parse_query("select a from t where foo = '5';");
	query q = parse_query("select value, time from t where c = '5';");
	printf("table name: %s\n", q.f->table);
	printf("where column: %s, value: %s\n", q.w->column_name, q.w->v->svalue);
	printf("arr: %d\n", q.c->size);
	int i = 0;
	for (i = 0; i < q.c->size; i++) {
		printf("%s\n", q.c->elems[i]);
	}
	close_query(&q);
	return 0;
}
